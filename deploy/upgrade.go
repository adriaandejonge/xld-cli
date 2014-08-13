package deploy

import (
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/intf"
)

var UpdateCmd cmd.Option = cmd.Option{
	Do:          update,
	Name:        "update",
	Description: "Updates existing deployables in a deployment",
	Permission: "deploy#upgrade",
	MinArgs:    2,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func update(args intf.Command) (result string, err error) {
	return execute(args, "UPDATE")
}
