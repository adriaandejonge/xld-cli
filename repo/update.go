package repo

import (
	"errors"
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/intf"
)

var UpdateCmd cmd.Option = cmd.Option{
	Do:          update,
	Name:        "update",
	Description: "Change existing configuration item",
	Permission:  "repo#edit",
	MinArgs:     0,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func update(args intf.Command) (result string, err error) {
	return "error", errors.New("xld update not yet implemented")
}
