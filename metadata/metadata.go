package metadata

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/adriaandejonge/xld/http"
	"github.com/adriaandejonge/xld/intf"
)

type (
	List struct {
		CITypes []CIType `xml:"descriptor"`
	}

	CIType struct {
		Type         string        `xml:"type,attr"`
		Virtual      bool          `xml:"virtual,attr"`
		Versioned    bool          `xml:"versioned,attr"`
		Root	     string        `xml:"root,attr"`
		Properties   []Property    `xml:"property-descriptors>property-descriptor"`
		ControlTasks []ControlTask `xml:"control-tasks>control-task"`
		Interfaces   []string      `xml:"interfaces>interface"`
		SuperTypes   []string      `xml:"superTypes>superType"`
		props map[string]Property
	}

	Property struct {
		Name               string `xml:"name,attr"`
		Fqn                string `xml:"fqn,attr"`
		Label              string `xml:"label,attr"`
		Kind               string `xml:"kind,attr"`
		Description        string `xml:"description,attr"`
		Category           string `xml:"category,attr"`
		AsContainment      bool   `xml:"asContainment,attr"`
		Inspection         bool   `xml:"inspection,attr"`
		Required           bool   `xml:"required,attr"`
		RequiredInspection bool   `xml:"requiredInspection,attr"`
		Password           bool   `xml:"password,attr"`
		Transient          bool   `xml:"transient,attr"`
		Size               string `xml:"size,attr"`
		Default            string `xml:"default,attr"`
		Hidden             bool   `xml:"hidden,attr"`
		ReferencedType     string `xml:"referencedType"`
	}

	ControlTask struct {
		Name        string `xml:"name,attr"`
		Fqn         string `xml:"fqn,attr"`
		Label       string `xml:"label,attr"`
		Description string `xml:"description,attr"`
	}
)

var shorthand = map[string]string{
	"dict": "udm.Dictionary",
	"dir": "core.Directory",
	"env": "udm.Environment",
}

func (ciType *CIType) Prop(name string) Property {
	if ciType.props == nil {
		ciType.indexProps()
	}
	return ciType.props[name]
}

func (ciType *CIType) indexProps() {
	ciType.props = make(map[string]Property)
	for _, prop := range ciType.Properties {
		ciType.props[prop.Name] = prop
	}
}

func Do(args intf.Command) (result string, err error) {

	subs := args.Subs()
	switch args.Main() {
	case "types":
		// TODO check nr of args again
		return types()
	case "describe":
		if len(subs) == 0 {
			return "error", errors.New("xld metadata expects at least 1 argument")
		} 
		
		for _, sub := range subs {
			_, err := describe(sub)
			if err != nil {
				return "error", err
			}
		}
		return "", nil
		 

	// TODO orchestrators
	// TODO permissions
	default:
		return "error", errors.New("Unknown command")
	}

}
func types() (result string, err error) {
	_, body, err := http.Read("/metadata/type")
	if err != nil {
		return
	}
	// TODO check statuscode

	list := List{}
	err = xml.Unmarshal(body, &list)
	if err != nil {
		//fmt.Printf("error: %v", err)
		return
	}

	for _, ciType := range list.CITypes {
		fmt.Println(ciType.Type)
	}

	return "", nil

}
func describe(typeName string) (result string, err error) {
	ciType, err := Type(typeName)
	fmt.Println(ciType.Type + ":")

	for _, prop := range ciType.Properties {
		fmt.Println("  -" + prop.Name, prop.Kind, iif(prop.Required, "required", ""), iif(prop.Hidden, "hidden", ""))
	}

	return "", nil
}

func Type(typeName string) (retType *CIType, err error) {
	long := shorthand[typeName]
	if long != "" {
		typeName = long
	}
	//fmt.Println("Typename = ", typeName)

	_, body, err := http.Read("/metadata/type/" + typeName)
	if err != nil {
		return
	}
	// TODO check statuscode

	ciType := CIType{}
	err = xml.Unmarshal(body, &ciType)
	if err != nil {
		return
	}
	return &ciType, nil
}

// TODO: replace with method that works on the instance
func iif(cond bool, iftrue interface{}, iffalse interface{}) interface{} {
	if cond {
		return iftrue
	} else {
		return iffalse
	}
}
