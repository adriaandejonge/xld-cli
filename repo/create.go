package repo

import (
	"bytes"
	"errors"
	"github.com/adriaandejonge/xld/metadata"
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
	"github.com/clbanning/mxj/j2x"
	"strings"
)

var CreateCmd cmd.Option = cmd.Option{
	Do:          create,
	Name:        "create",
	Description: "Create new configuration item",
	Permission:  "repo#edit",
	MinArgs:     0,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func create(args intf.Command) (result string, err error) {
	subs := args.Subs()
	typeName := subs[0]
	ciName := subs[1]

	ciType, err := metadata.Type(typeName)
	if err != nil {
		return
	}

	// put this as the root in a map containing a map
	// do this AFTER the for loop

	// create new map and put the below in it

	mapProps := make(map[string]interface{})

	//props := args[2:]
	props := args.Arguments()
	for _, prop := range props {
		key := prop.Name()

		kind := ciType.Prop(key).Kind

		if kind == "" {
			return "error", errors.New("Unknown property type " + ciType.Type + "->" + key)
		}

		switch kind {

		case "BOOLEAN", "INTEGER", "STRING", "ENUM":
			mapProps[key] = prop.Value()

		case "CI":
			mapProps[key] = mapRef(prop.Value())

		case "MAP_STRING_STRING":
			mapProps[key] = mapStringString(prop.Map())

		case "SET_OF_STRING", "LIST_OF_STRING":
			mapProps[key] = mapSetOfStrings(prop.Values())

		case "SET_OF_CI", "LIST_OF_CI":
			mapProps[key] = mapSetOfCis(prop.Values())

		default:
			return "error", errors.New("Unknown property kind " + kind + " --> XLD server newer than client?")

		}
	}

	id := ciName
	if ciType.Root != "" {
		id = ciType.Root + "/" + id
	}
	id = AntiAbbreviate(id)
	mapProps["-id"] = id

	final := map[string]interface{}{ciType.Type: mapProps}

	// TODO Make this a util?
	json, _ := j2x.MapToJson(final)
	xml, _ := j2x.JsonToXml(json)

	body, err := http.Create("/repository/ci/"+id, bytes.NewBuffer(xml))

	return string(body), err
}

func mapStringString(kvPairs map[string]string) interface{} {
	entry := make([]map[string]interface{}, 0)

	for k, v := range kvPairs {
		entry = append(entry, map[string]interface{}{"-key": k, "#text": v})
	}
	return map[string]interface{}{"entry": entry}
}

func mapSetOfStrings(values []string) interface{} {
	return map[string]interface{}{"value": values}
}

func mapSetOfCis(values []string) interface{} {
	cis := make([]map[string]interface{}, 0)

	for _, ref := range values {
		cis = append(cis, mapRef(strings.TrimSpace(ref)))
	}
	return map[string]interface{}{"ci": cis}

}
func mapRef(value string) map[string]interface{} {
	// TODO read @ROOT for type of ref
	// TODO or provide default for virtual type

	return map[string]interface{}{"-ref": AntiAbbreviate(value)}
}
