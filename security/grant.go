package security

import (
 	"fmt"
	_ "strings"
	_"encoding/xml"
	"strconv"
	_ "errors"

	"github.com/adriaandejonge/xld/repo"

	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"

	"github.com/clbanning/mxj/x2j"
	"github.com/clbanning/mxj/j2x"

	_ "reflect"
	_ "bytes"
)

var GrantCmd cmd.Option = cmd.Option{
	Do:          grant,
	Name:        "grant",
	Description: "Add permission to a directory or global",

	Permission: "security#edit",
	MinArgs:    3,
	Help: `
# xld grant <ci> <role> -<permission(s)>...

Add permissions for the specified role on the specified configuration item.

Usage:

 - xld grant <ci> <role> -<permission(s)>...

 Examples:

 - xld grant env/MyDir ops -deploy#initial
 - xld grant global dev -login
`,
}

func grant(args intf.Command) (result string, err error) {
	subs := args.Subs()
	
	ci := subs[0]
	roleToChange := subs[1]
	perms := subs[2:]

	// Get current permissions
	// (already works for global keyword :) )
	body, err := http.Read("/internal/security/roles/permissions/" + repo.AntiAbbreviate(ci))
	if err != nil {
		return
	}

	values, err := x2j.XmlToMap(body)
	if err != nil {
		return
	}

	arr := arrayFromMap(values["collection"], "rolePermissions")
	changed := false

	for _, elMap := range arr {
		role := elMap["role"].(map[string]interface{})
		if role["-name"] == roleToChange {

			oldPerms := elMap["permissions"].([]interface{})
			for _, el := range perms {
				oldPerms = append(oldPerms, el)	
			}
			elMap["permissions"], changed = oldPerms, true
		}
	}

	if !changed {
		roleMap := map[string]string{"-id": strconv.Itoa(len(arr)), "-role": roleToChange}
		fmt.Println("Need to add it.... :( ", roleMap)
		fullMap := map[string]interface{}{"permissions": perms, "role": roleMap}
		arr = append(arr, fullMap)
		values["collection"].(map[string]interface{})["rolePermissions"] = arr
	}
	
	// TODO Make this a util?
	json, _ := j2x.MapToJson(values)
	xml, _ := j2x.JsonToXml(json)

	fmt.Println("modified(?) response = ", string(xml))

	/*
	body, err = http.Update("/internal/security/roles/permissions/" + 
		repo.AntiAbbreviate(ci), 
		bytes.NewBuffer(xml))
	if err != nil {
		return
	}

	fmt.Println("modified(?) response = ", body)
	//*/



	
	//fmt.Println("permset = ", permSet)


	// decode XML to ... JSON or marshall to objects?

	// is role already in there?
	// if not: add role

	// is permission already in there? -> ignore

	// add permission

	// encode as XML

	// post to http
	return
}

func arrayFromMap(intfMap interface{}, key string) []map[string]interface{} {
	rMap, ok := intfMap.(map[string]interface{})
	if !ok {
		//return "error", errors.New("COLLECTIONS = EMPTY")
		rMap = make(map[string]interface{})
		return make([]map[string]interface{}, 0)
	}

	arr, ok := rMap["rolePermissions"].([]interface{})
	if !ok {
		return make([]map[string]interface{}, 0)	
	}

	arrMap := make([]map[string]interface{}, len(arr))

	for i, el := range arr {
		arrMap[i], ok = el.(map[string]interface{})
		if !ok {
			return make([]map[string]interface{}, 0)	
		}
	}
	fmt.Println("6", arrMap)

	return arrMap
}