package deploy

import (
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/intf"
)

var DeployCmd cmd.Option = cmd.Option{
	Do:          deploy,
	Name:        "deploy",
	Description: "Execute a deployment",
	Permission: "deploy#initial",
	MinArgs:    2,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func deploy(args intf.Command) (result string, err error) {
	return execute(args, "INITIAL")
}
