package repo

import (
	"encoding/json"
	"errors"
	"github.com/adriaandejonge/xld/metadata"
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
	"github.com/clbanning/mxj/x2j"
	"strings"
)

var ReadCmd cmd.Option = cmd.Option{
	Do:          read,
	Name:        "read",
	Description: "Read configuration item",
	Permission:  "read",
	MinArgs:     1,
	Help: `
# xld read <id>

Read a configuraton item from the repository and output JSON format.

Usage:

- xld read <id>

Examples:

- xld read env/MyEnv
- xld read inf/MyServer/MyTomcat

Note: env and inf are abbreviations for Environments and Infrastructure. You can also use the full names:

- xld read Infrastructure/MyServer

You can also use the abbreviation "latest" to automatically find the newest version of an application:

- xld read app/MyApp/latest
`,
}

func read(args intf.Command) (result string, err error) {

	resultMap, err := ReadAsMap(args)
	if err != nil {
		return "error", err
	}
	delete(resultMap, "token")
	
	json, err := json.MarshalIndent(resultMap, "", "    ")
	if err != nil {
		return "error", err
	}

	return string(json), err

}

func ReadAsMap(args intf.Command) (resultMap map[string]interface{}, err error) {
	subs := args.Subs()

	ci := AntiAbbreviate(subs[0])

	body, err := http.Read("/repository/ci/" + ci)
	if err != nil {
		return
	}

	return xml2CleanMap(body)

}

func Xml2CleanJson(body []byte) (cleanJson string, err error) {
	myMap, err := xml2CleanMap(body)
	if err != nil {
		return "error", err
	}
	delete(myMap, "token")

	json, err := json.MarshalIndent(myMap, "", "    ")
	if err != nil {
		return "error", err
	}

	return string(json), err

}

func xml2CleanMap(body []byte) (resultMap map[string]interface{}, err error) {
	values, err := x2j.XmlToMap(body)
	if err != nil {
		return
	}

	cleanProperties := make(map[string]interface{})
	resultMap = make(map[string]interface{})

	for key, value := range values {
		resultMap["type"] = key

		ciType, err := metadata.Type(key)
		if err != nil {
			return nil, err
		}

		valueMap := value.(map[string]interface{})

		for k, v := range valueMap {
			if strings.HasPrefix(k, "-") {

				resultMap[k[1:]] = v

			} else {
				cleanProperties[k], err = readProperty(k, v, ciType)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	resultMap["content"] = cleanProperties
	return resultMap, nil
}

func readProperty(key string, value interface{}, ciType *metadata.CIType) (result interface{}, err error) {
	kind := ciType.Prop(key).Kind

	switch kind {

	case "BOOLEAN", "INTEGER", "STRING", "ENUM":
		result = value.(string)

	case "CI":
		result = cleanRef(value)

	case "MAP_STRING_STRING":
		result = cleanStringString(value)

	case "SET_OF_STRING", "LIST_OF_STRING":
		result = cleanSetOfStrings(value)

	case "SET_OF_CI", "LIST_OF_CI":
		result = cleanSetOfCis(value)

	default:
		return "error", errors.New("Unknown property kind " + kind + " --> XLD server newer than client?")
	}

	return
}

func cleanStringString(input interface{}) map[string]string {
	empty := make(map[string]string)
	if input != nil {
		switch input.(type) {
		case map[string]interface{}:
			resultMap := make(map[string]string)

			rootMap := input.(map[string]interface{})
			entries := rootMap["entry"].([]interface{})
			for _, el := range entries {
				keyVal := el.(map[string]interface{})
				resultMap[keyVal["-key"].(string)] = keyVal["#text"].(string)
			}

			return resultMap
		default:
			return empty
		}
	} else {
		return empty
	}
}

func cleanSetOfStrings(input interface{}) []string {
	empty := make([]string, 0)
	if input != nil {
		switch input.(type) {
		case map[string]interface{}:
			valuesMap := input.(map[string]interface{})
			valuesArr := valuesMap["value"]
			switch valuesArr.(type) {
			case string:
				return []string{valuesArr.(string)}
			case []interface{}:
				intfArr := valuesArr.([]interface{})
				stringArr := make([]string, len(intfArr))
				for i, el := range intfArr {
					stringArr[i] = el.(string)
				}
				return stringArr
			default:
				return empty
			}
		default:
			return empty
		}
	} else {
		return empty
	}
}

func cleanSetOfCis(input interface{}) []string {
	empty := make([]string, 0)
	if input != nil {
		switch input.(type) {
		case map[string]interface{}:

			ciMapsIf := input.(map[string]interface{})["ci"]

			ciMaps := arrayOfMaps(ciMapsIf)

			resultArr := make([]string, len(ciMaps))

			for i, ciMap := range ciMaps {
				ref := ciMap.(map[string]interface{})["-ref"]
				resultArr[i] = ref.(string)
			}

			return resultArr
		default:
			return empty
		}
	} else {
		return empty
	}
}

func arrayOfMaps(input interface{}) []interface{} {
	empty := make([]interface{}, 0)
	if input != nil {
		switch input.(type) {
		case map[string]interface{}:
			return []interface{}{input}
		case []interface{}:
			return input.([]interface{})
		default:
			return empty
		}
	} else {
		return empty
	}
}

func cleanRef(input interface{}) string {
	empty := ""
	if input != nil {
		switch input.(type) {
		case map[string]interface{}:
			ciMap := input.(map[string]interface{})
			ref := ciMap["-ref"]
			return ref.(string)
		default:
			return empty
		}
	} else {
		return empty
	}
}
