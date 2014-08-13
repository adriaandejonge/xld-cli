package repo

import (
	"errors"
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/intf"
)

var ModifyCmd cmd.Option = cmd.Option{
	Do:          modify,
	Name:        "modify",
	Description: "Change existing configuration item",
	Permission:  "repo#edit",
	MinArgs:     0,
	Help: `
# XLD Modify: 

Not yet implemented
`,
}

func modify(args intf.Command) (result string, err error) {
	return "error", errors.New("xld update not yet implemented")
}
