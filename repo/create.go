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
	MinArgs:     1,
	Help: `
# XLD Create: 

Create items in XL Deploy from command line.

## Basic usage:

xld create <type> <id> -<key> <value(s)>...

## Advanced usage:

 - To enter key-value pairs, you can pipe JSON or CSV as input:

	<output json map> | xld create <type> <id> -<key> stdin:json

	<ouput csv file> | xld create <type> <id> -<key> stdin:csv

 - To enter the full content, you can pipe JSON:

	<output json map> | xld create <type> <id> stdin

 - To enter the full content, type and ID, you can pipe JSON:

	<output json map> | xld create stdin

Examples:

xld create overthere.LocalHost inf/MyServer -os UNIX -tags one two three -temporaryDirectoryPath /tmp

xld create dict env/MyDict -entries key1=value1 key2=value2

xld create env env/MyEnv -members inf/MyServer -dictionaries env/MyDict


Take a file myentries.json with the following content:

{
	"key1": "value1",
	"key2": "value2"
}

and type:

cat myentries.json | xld create dict env/MyDict -entries stdin:json

Take a file mydict.json with the following content:

{
	"entries": {
		"key1": "value1",
		"key2": "value2"
	}
}

and type:

cat myentries.json | xld create dict env/MyDict stdin:json

Take a file myitem.json with the following content:

{
    "content": {
        "entries": {
            "key1": "value1",
            "key2": "value2"
        }
    },
    "id": "env/MyDict",
    "type": "dict"
}

and type:

cat myentries.json | xld create stdin:json

Abbreviations

XLD allows the following abbreviations for item types:

env -> udm.Environment
dict -> udm.Dictionary
dir -> udm.Directory

XLD allows the following abbreviations for ID paths:

app -> Applications
env -> Environments
inf -> Infrastructure
conf -> Configuration

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
	id = AntiAbbreviate(id)
	if ciType.Root != "" && !strings.HasPrefix(id, ciType.Root) {
		id = ciType.Root + "/" + id
	}
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
