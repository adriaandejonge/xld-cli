package repo

import (
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
)

var RemoveCmd cmd.Option = cmd.Option{
	Do:          remove,
	Name:        "remove",
	Description: "Remove existing configuration item",
	Permission:  "import#remove", //repo#edit
	MinArgs:     0,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func remove(args intf.Command) (result string, err error) {
	subs := args.Subs()
	ciName := AntiAbbreviate(subs[0])
	// TODO validate input

	body, err := http.Remove("/repository/ci/" + ciName)

	result = string(body)

	return
}
