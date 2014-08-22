package repo

import (
	"encoding/json"
	"errors"
	"strings"
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/intf"
	"github.com/adriaandejonge/xld/metadata"
)

var ModifyCmd cmd.Option = cmd.Option{
	Do:          modify,
	Name:        "modify",
	Description: "Change existing configuration item",
	Permission:  "repo#edit",
	MinArgs:     0,
	Help: `
# xld modify <id> -<key> <value(s)>...

Take an existing configuration item and modify one or more fields. Only specify the fields that need to be modified.

Usage:

 - xld modify <id> -<key> <value(s)>...
 - xld modify <id> -add-<key> <value(s)>...
 - xld modify <id> -remove-<key> <value(s)>...
 - xld modify <id> -change-<key> <value(s)>...

The add/change/remove prefixes are helpers for modifying lists and maps without respecifying the full list or map again.

Examples:

 - xld modify env/MyDict -entries key1=value1 key2=newValue2
 - xld modify env/MyDict -add-entries key3=value3
 - xld modify env/MyDict -remove-entries key1
 - xld modify env/MyDict -change-entries key2=newValue2
`,
}

func modify(args intf.Command) (result string, err error) {
	resultMap, err := ReadAsMap(args)
	if err != nil {
		return "error", err
	}

	content := resultMap["content"]
	mapContent, ok := content.(map[string]interface{})
	if !ok {
		return "error", errors.New("Unable to modify existing content")
	}

	props := args.Arguments()
	for _, prop := range props {
		key := prop.Name()

		split := strings.SplitN(key, "-", 2)
		prefix := split[0]
		switch prefix {
		case "add", "remove", "change":
			if len(split) > 1 {
				key = split[1]
			}
			err = handleAddRemoveChange(mapContent, prefix, key, prop)
			if err != nil {
				return "error", err
			}
		default:
			typeName, ok := resultMap["type"].(string)
			if !ok {
				return "error", errors.New("Cannot resolve type for existing item")
			}

			ciType, err := metadata.Type(typeName)
			if err != nil {
				return "error", err
			}

			err = handleDefault(mapContent, key, prop, ciType)
			if err != nil {
				return "error", err
			}
		}
	}

	json, err := json.MarshalIndent(resultMap, "", "    ")

	createCmd, err := cmd.NewStdinCmd("modify-sub", string(json))
	if err != nil {
		return "error", err
	}
	return createOrModify(createCmd)
}

func handleDefault(mapContent map[string]interface{}, key string, prop intf.Argument, ciType *metadata.CIType) error {
	kind := ciType.Prop(key).Kind

	switch kind {
	case "MAP_STRING_STRING":
		mapContent[key] = prop.Map()
	case "SET_OF_STRING", "LIST_OF_STRING", "SET_OF_CI", "LIST_OF_CI":
		mapContent[key] = prop.Values()
	case "BOOLEAN", "INTEGER", "STRING", "ENUM", "CI":
		mapContent[key] = prop.Value()
	default:
		return errors.New("Unexpected value type read from repository; does it exist?")
	}
	return nil
}

func handleAddRemoveChange(mapContent map[string]interface{}, prefix string, key string, prop intf.Argument) error {
	val := mapContent[key]

	// ASSUMPTION: maps and lists are never empty in xld read
	switch t := val.(type) {
	case map[string]string:

		switch prefix {
		case "add", "change":
			for k, v := range prop.Map() {
				t[k] = v
			}
		case "remove":
			for _, el := range prop.Values() {
				delete(t, el)
			}

		}
	case []string:
		switch prefix {
		case "add", "change":
			for _, el := range prop.Values() {
				t = append(t, el)
			}
			mapContent[key] = t
		case "remove":
			newElements := make([]string, 0)
			for _, el := range t {
				found := false
				for _, cmdArg := range prop.Values() {
					if el == cmdArg {
						found = true
					}
				}
				if !found {
					newElements = append(newElements, el)
				}
			}
			mapContent[key] = newElements			
		}
	default:
		return errors.New("Unexpected value type read from repository; does it exist?")
	}
	return nil

}
