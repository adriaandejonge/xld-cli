package metadata

import (
	"encoding/xml"
	"errors"
	"github.com/adriaandejonge/xld/util/http"
	"io"
)

func Type(typeName string) (retType *CIType, err error) {
	long := shorthand[typeName]
	if long != "" {
		typeName = long
	}

	body, err := http.Read("/metadata/type/" + typeName)
	if err != nil {
		return
	}
	// TODO check statuscode

	ciType := CIType{}
	err = xml.Unmarshal(body, &ciType)

	if err == io.EOF {
		return nil, errors.New("Type " + typeName + " not found")
	} else if err != nil {
		return
	}

	return &ciType, nil
}

var shorthand = map[string]string{
	"dict": "udm.Dictionary",
	"dir":  "core.Directory",
	"env":  "udm.Environment",
}

// TODO: replace with method that works on the instance
func iif(cond bool, iftrue interface{}, iffalse interface{}) interface{} {
	if cond {
		return iftrue
	} else {
		return iffalse
	}
}
