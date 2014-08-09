package repo

import (
	"fmt"
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
	"net/url"
	"strings"
)

var ListCmd cmd.Option = cmd.Option{
	Do:          list,
	Name:        "list",
	Description: "List configuration items",
	Permission:  "",
	MinArgs:     0,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func list(args intf.Command) (result string, err error) {
	subs := args.Subs()

	arguments := make([]string, 0)

	if len(subs) > 0 {
		if strings.HasSuffix(subs[0], "*") {
			arguments = append(arguments, "ancestor="+url.QueryEscape(AntiAbbreviate(strings.Replace(subs[0], "*", "", -1))))
		} else {
			arguments = append(arguments, "parent="+url.QueryEscape(AntiAbbreviate(subs[0])))
		}
	}

	extra := args.Arguments()
	for _, el := range extra {
		name := el.Name()

		switch name {
		case "type":
			arguments = append(arguments, "type="+url.QueryEscape(el.Value()))
		case "like":
			arguments = append(arguments, "namePattern="+url.QueryEscape(el.Value()))
		case "before":
			//TODO lastModifiedBefore
		case "after":
			//TODO lastModifiedAfter
		case "page":
			arguments = append(arguments, "page="+url.QueryEscape(el.Value()))
		case "pagesize":
			arguments = append(arguments, "resultsPerPage="+url.QueryEscape(el.Value()))
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
