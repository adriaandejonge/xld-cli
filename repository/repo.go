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

	ciType, _ := metadata.Type(typeName)
	// TODO check for error



	// put this as the root in a map containing a map
	// do this AFTER the for loop

	// create new map and put the below in it

	mapProps := make(map[string]interface{})

	props := args[2:]
	for _, prop := range props {
		keyval := strings.SplitN(prop, "=", 2)
		key := keyval[0]
		value := keyval[1]
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
		fmt.Println("prop = ", ciType.Prop(key))

		// SWITCH ciType.Prop(key) - Type

		mapProps[key] = value

		// let map entry depend on type
		// switch? or OOP? (suggest start with switch)

	}

	// create XML based on map

	// TODO: Clean up
	fmt.Println("Props = ", mapProps)
	fmt.Println("ciType =", ciType)
	fmt.Println("arg 1 =", args[1])
	fmt.Println("arg 2 =", args[2])

	return
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
