package repo

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"github.com/adriaandejonge/xld/metadata"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
	"github.com/clbanning/mxj/j2x"
	"strings"
	"net/url"
)

var shorthand = map[string]string{
	"app":  "Applications",
	"env":  "Environments",
	"inf":  "Infrastructure",
	"conf": "Configuration",
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
		//key, value := keyValue(prop, "=")
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

func list(args intf.Command) (result string, err error) {
	//http://localhost:4516/deployit/repository/query?ancestor=Environments

	subs := args.Subs()

	arguments := make([]string, 0)

	if len(subs) > 0 {
		if strings.HasSuffix(subs[0], "*") {
			arguments = append(arguments, "ancestor=" + url.QueryEscape(AntiAbbreviate(strings.Replace(subs[0], "*", "", -1))))
		} else {
			arguments = append(arguments, "parent=" + url.QueryEscape(AntiAbbreviate(subs[0])))
		}
	}

	extra := args.Arguments()
	for _, el := range extra {
		name := el.Name()

		switch name {
		case "type":
			arguments = append(arguments, "type=" + url.QueryEscape(el.Value()))
		case "like":
			arguments = append(arguments, "namePattern=" + url.QueryEscape(el.Value()))
		case "before":
			//TODO lastModifiedBefore	
		case "after":
			//TODO lastModifiedAfter
		case "page":
			arguments = append(arguments, "page=" + url.QueryEscape(el.Value()))
		case "pagesize":
			arguments = append(arguments, "resultsPerPage=" + url.QueryEscape(el.Value()))
		}
	}


	body, err := http.Read("/repository/query?" + strings.Join(arguments, "&"))
	list, err := readCiList(body)
	if err != nil {
		return "error", err
	}
	for _, el := range list.CIs {
		fmt.Println(el.CiRef)
	}



	return "", err
}


type (
	List struct {
		CIs         []Ci `xml:"ci"`
	}

	Ci struct {
		CiRef         string    `xml:"ref,attr"`
		Type            string `xml:"type,attr"`
	}
)


type ByVersion [][]string

func (versions ByVersion) Len() int {
    return len(versions)
}


// Simple heuristic to avoid converting strings and integers
// Please mind: A goes before AA and B goes before AA.
// In other words, shorter strings go before longer.

// Alternative:
	//matcher := regexp.MustCompile("^[0-9]+$")
	//sort
	//	matcher.MatchString(el) -> true / false
	// -> decide string vs int compare

// TODO: if necessary, speed this up; now inefficient but not noticable

const norm string = "0000000000"

func normalizeLess(a string, b string) bool {
	return (norm + a)[len(a):] < (norm + b)[len(b):]
}

func normalizeMore(a string, b string) bool {
	return (norm + a)[len(a):] > (norm + b)[len(b):]
}

func less(versions ByVersion, index, i, j int) bool {
	if normalizeLess(versions[i][index], versions[j][index]) {
		return true
	} else if normalizeMore(versions[i][index], versions[j][index]) {
		return false
	} else if index < len(versions[i]) - 1 {
		return less(versions, index+1, i, j)
	} else {
		return false
	}
}

func (versions ByVersion) Less(i, j int) bool {
	return less(versions, 0, i, j)
}
func (versions ByVersion) Swap(i, j int) {
    versions[i], versions[j] = versions[j], versions[i]
}


func readCiList(body []byte) (list List, err error) {
	list = List{}
	err = xml.Unmarshal(body, &list)
	return
}

func update(args intf.Command) (result string, err error) {
	return "error", errors.New("xld update not yet implemented")
}

func remove(args intf.Command) (result string, err error) {
	subs := args.Subs()
	ciName := AntiAbbreviate(subs[0])
	// TODO validate input

	body, err := http.Remove("/repository/ci/" + ciName)

	result = string(body)

	return
}

// TODO Make this a util?
func AntiAbbreviate(ciName string) string {
	prefix := strings.SplitN(ciName, "/", 2)
	longer := shorthand[prefix[0]]

	remainder := ""

	if len(prefix) > 1 {
		remainder = prefix[1]

	}

	if longer != "" {		
		ciName = longer + "/" + remainder
	}

	if strings.HasSuffix(ciName, "latest") {
		searchFor := ciName[:len(ciName) - len("latest")]
		body, err := http.Read("/repository/query?parent=" + searchFor)
		list, err := readCiList(body)
		if err != nil {
			// TODO ERROR
			return "error"
		}

		splitter := regexp.MustCompile("[\\.?\\-?_?]+")

		versions := make([][]string, 0)
		ciMap := make(map[string]Ci)
		for _, el := range list.CIs {
			version := el.CiRef[len(searchFor):]
			
			split := splitter.Split(version, -1)
			ciMap[strings.Join(split, ".")] = el
			versions = append(versions, split)
		}
		sort.Sort(ByVersion(versions))
		ciName = ciMap[strings.Join(versions[len(versions) - 1], ".")].CiRef

	}

	return ciName
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

func keyValue(combined string, split string) (key string, value string) {
	keyval := strings.SplitN(combined, split, 2)
	return keyval[0], keyval[1]

}
