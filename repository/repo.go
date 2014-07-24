package repository

import (
	"errors"
	"fmt"
	"strings"
	"bytes"
	"github.com/adriaandejonge/xld/http"
	"github.com/adriaandejonge/xld/metadata"
	"github.com/clbanning/mxj/j2x"
)

var shorthand = map[string]string{
	"app": "Applications",
	"env": "Environments",
	"inf": "Infrastructure",
	"conf": "Configuration",
}


func Do(args []string) (result string, err error) {

	if len(args) < 1 {
		return "error", errors.New("xld repo expects at least 1 argument")
	} else {
		if err != nil {
			return "error", err
		}

		switch args[0] {
		case "create":
			return create(args[1:])

		case "remove":
			return remove(args[1:])

		default:
			return "error", errors.New("Unknown command")
		}

	}

}

func create(args []string) (result string, err error) {
	typeName := args[0]
	ciName := args[1]

	ciType, err := metadata.Type(typeName)
	if err != nil {
		return
	}
	



	// put this as the root in a map containing a map
	// do this AFTER the for loop

	// create new map and put the below in it

	mapProps := make(map[string]interface{})

	props := args[2:]
	for _, prop := range props {
		key, value := keyValue(prop, "=")

		kind := ciType.Prop(key).Kind

		if kind == "" {
			return "error", errors.New("Unknown property type " + ciType.Type + "->" + key)
		}

		switch kind {

		case "BOOLEAN", "INTEGER", "STRING", "ENUM":
			mapProps[key] = value
		case "CI":
			mapProps[key] = mapRef(value)
			
			
		case "MAP_STRING_STRING":
			entry := make([]map[string]interface{}, 0)

			kvPairs := strings.Split(value, " ")
			for _, kvPair := range kvPairs {
				k, v := keyValue(kvPair, ":")
				entry = append(entry, map[string]interface{}{"-key": k, "#text": v})
			}
			mapProps[key] = map[string]interface{}{"entry": entry}
		case "SET_OF_STRING", "LIST_OF_STRING":
			values := strings.Split(value, ",")
			// TODO filter spaces
			// $.map() like function??? (wbn)
			mapProps[key] = map[string]interface{}{"value": values}
		case "SET_OF_CI", "LIST_OF_CI":
			mapProps[key] = mapSetOfCis(value)

		default:
			return "error", errors.New("Unknown property kind " + kind + " --> XLD server newer than client?")
			
		}
	}

	id := ciName
	if ciType.Root != "" {
		id = ciType.Root + "/" + id
	}
	mapProps["-id"] = id

	final := map[string]interface{}{ciType.Type: mapProps}




	json, _ := j2x.MapToJson(final)
	xml, _ := j2x.JsonToXml(json)


	statusCode, body, err := http.Create("/repository/ci/" + id, bytes.NewBuffer(xml))

	if statusCode != 200 {
		err = errors.New(fmt.Sprintf("HTTP status code %d: %s", statusCode, body))
		// TODO if message type is XML (validation-message), then read and display nicely
	}

	return
}

func remove(args []string) (result string, err error) {
	ciName := antiAbbreviate(args[0])

	statusCode, body, err := http.Delete("/repository/ci/" + ciName)

	result = string(body)


	if statusCode < 200 || statusCode >= 300 {
		err = errors.New(fmt.Sprintf("HTTP status code %d: %s", statusCode, body))
	}

	return 
}

func antiAbbreviate(ciName string) string {
	prefix := strings.SplitN(ciName, "/", 2)
	longer := shorthand[prefix[0]]

	if longer != "" {
		ciName = longer + "/" + prefix[1]
	}
	return ciName
}

func mapSetOfCis(value string) interface{} {
	cis := make([]map[string]interface{}, 0)

	ciRefs := strings.Split(value, ",")
	for _, ref := range ciRefs {
		cis = append(cis, mapRef(strings.TrimSpace(ref)))
	}
	return map[string]interface{}{"ci": cis}

}
func mapRef(value string) map[string]interface{} {
	// TODO read @ROOT for type of ref
	// TODO or provide default for virtual type

	return map[string]interface{}{"-ref": value}
}

func keyValue(combined string, split string) (key string, value string) {
	keyval := strings.SplitN(combined, split, 2)
	return keyval[0], keyval[1]

}
