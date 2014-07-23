package repository

import (
	"errors"
	"fmt"
	"strings"
	"github.com/adriaandejonge/xld/http"
	"github.com/adriaandejonge/xld/metadata"
	"github.com/clbanning/mxj"
	"github.com/clbanning/mxj/j2x"
)



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
			//return repo()
		default:
			return "error", errors.New("Unknown command")
		}

	}

}

func create(args []string) (result string, err error) {
	typeName := args[0]
	//ciName := args[1]

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
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
		fmt.Println("prop = ", ciType.Prop(key))

		kind := ciType.Prop(key).Kind

		switch kind {

		case "BOOLEAN", "INTEGER", "STRING", "ENUM":
			// TODO Check that this is correct
			mapProps[key] = value
		case "CI":
			mapProps[key] = map[string]interface{}{"-ref": value}
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
			cis := make([]map[string]interface{}, 0)

			ciRefs := strings.Split(value, ",")
			for _, ref := range ciRefs {
				cis = append(cis, map[string]interface{}{"-ref": ref})
			}
			mapProps[key] = map[string]interface{}{"ci": cis}

		default:
			// Should not get here
			//return "error", errors.New("Unknown property type " + kind)
			return "error", errors.New("Unknown property type " + ciType.Type + "->" + key)
			
		}
	}

	final := map[string]interface{}{"ciType.Type": mapProps}


	json, _ := j2x.MapToJson(final)
	xml, _ := j2x.JsonToXml(json)

	// TODO Clean Up:
	fmt.Println("\n\n")
	fmt.Println("XML = ", string(xml))

	return
}

func keyValue(combined string, split string) (key string, value string) {
	keyval := strings.SplitN(combined, split, 2)
	return keyval[0], keyval[1]

}

// TODO Clean up if code is not needed as example anymore
func repo() (result string, err error) {

	_, body, err := http.Read("/repository/ci/Infrastructure/test")
	if err != nil {
		return
	}

	m, err := mxj.NewMapXml(body)
	if err != nil {
		return
	}

	return fmt.Sprint("Map = ", m), nil

}

// TODO Clean up if code is not needed as example anymore
func DoSomething() {

	myMap := map[string]interface{}{"test": map[string]interface{}{"key": "value", "keys": "values", "map": map[string]interface{}{"array": []string{"a", "b", "c"}}}}
	json, _ := j2x.MapToJson(myMap)
	xml, _ := j2x.JsonToXml(json)
	fmt.Println(string(xml))

}
