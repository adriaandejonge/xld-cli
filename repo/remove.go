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
	MinArgs:     1,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func remove(args intf.Command) (result string, err error) {

	for _, sub := range args.Subs() {
		ciName := AntiAbbreviate(sub)
		// TODO validate input

		body, err := http.Remove("/repository/ci/" + ciName)
		if err != nil {
			return "error", err
		}

		result = string(body)
	}

	return
}
