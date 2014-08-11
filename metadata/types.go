package metadata

import (
	"encoding/xml"
	"fmt"
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
)

var TypesCmd cmd.Option = cmd.Option{
	Do:          types,
	Name:        "types",
	Description: "List configuration types",
	Permission: "",
	MinArgs:    0,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func types(args intf.Command) (result string, err error) {
	body, err := http.Read("/metadata/type")
	if err != nil {
		return
	}
	
	list := List{}
	err = xml.Unmarshal(body, &list)
	if err != nil {
		return
	}

	for _, ciType := range list.CITypes {
		fmt.Println(ciType.Type)
	}

	return "", nil

}
