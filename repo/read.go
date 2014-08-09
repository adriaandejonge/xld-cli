package repo

import (
	"errors"
	"fmt"
	"github.com/adriaandejonge/xld/metadata"
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
	"github.com/clbanning/mxj/j2x"
	"github.com/clbanning/mxj/x2j"
	"strings"
)

var ReadCmd cmd.Option = cmd.Option{
	Do:          read,
	Name:        "read",
	Description: "Read configuration item",
	Help: `
TODO: 
	Long, multi-line help text
`,
	Permission: "read",
	MinArgs:    1,
}

func read(args intf.Command) (result string, err error) {
	subs := args.Subs()

	ci := AntiAbbreviate(subs[0])

	body, err := http.Read("/repository/ci/" + ci)

	values, err := x2j.XmlToMap(body)
	cleanProperties := make(map[string]interface{})
	resultMap := make(map[string]interface{})

	for key, value := range values {
		resultMap["type"] = key

		ciType, errr := metadata.Type(key)
		if errr != nil {
			// TODO why errr? (shadowed err)
			return "err", errr
		}

		valueMap := value.(map[string]interface{})

		for k, v := range valueMap {

			if strings.HasPrefix(k, "-") {

				resultMap[k[1:]] = v

			} else {

				kind := ciType.Prop(k).Kind

				switch kind {

				case "BOOLEAN", "INTEGER", "STRING", "ENUM":
					cleanProperties[k] = v.(string)

				case "CI":
					switch v.(type) {
					case map[string]interface{}:
						ciMap := v.(map[string]interface{})
						ref := ciMap["-ref"]
						cleanProperties[k] = ref.(string)
					default:
						cleanProperties[k] = ""
					}

				case "MAP_STRING_STRING":
					switch v.(type) {
					case map[string]interface{}:
						resultMap := make(map[string]string)

						rootMap := v.(map[string]interface{})
						entries := rootMap["entry"].([]interface{})
						for _, el := range entries {
							keyVal := el.(map[string]interface{})
							resultMap[keyVal["-key"].(string)] = keyVal["#text"].(string)
						}

						cleanProperties[k] = resultMap
					default:
						cleanProperties[k] = make(map[string]string)
					}

				case "SET_OF_STRING", "LIST_OF_STRING":
					switch v.(type) {
					case map[string]interface{}:
						valuesMap := v.(map[string]interface{})
						valuesArr := valuesMap["value"]
						switch valuesArr.(type) {
						case string:
							cleanProperties[k] = []string{valuesArr.(string)}
						case []interface{}:
							intfArr := valuesArr.([]interface{})
							stringArr := make([]string, len(intfArr))
							for i, el := range intfArr {
								stringArr[i] = el.(string)
							}
							cleanProperties[k] = stringArr
						default:
							// TODO Throw error?
							cleanProperties[k] = make([]string, 0)
							fmt.Println("   Key/val = ", k, v, kind)

						}
					default:
						cleanProperties[k] = make([]string, 0)
						fmt.Println("   Key/val = ", k, v, kind)

					}

				case "SET_OF_CI", "LIST_OF_CI":
					if v != nil {
						switch v.(type) {
						case map[string]interface{}:

							ciMapsIf := v.(map[string]interface{})["ci"]
							ciMaps := ciMapsIf.([]interface{})

							resultArr := make([]string, len(ciMaps))

							for i, ciMap := range ciMaps {
								ref := ciMap.(map[string]interface{})["-ref"]
								resultArr[i] = ref.(string)
							}

							cleanProperties[k] = resultArr
						default:
							cleanProperties[k] = make([]string, 0)
						}
					}

				default:
					return "error", errors.New("Unknown property kind " + kind + " --> XLD server newer than client?")

				}

			}

		}
	}

	resultMap["props"] = cleanProperties
	json, _ := j2x.MapToJson(resultMap)

	return string(json), err

}
