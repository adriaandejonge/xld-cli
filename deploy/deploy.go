package deploy

import (
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/intf"
)

var InitialCmd cmd.Option = cmd.Option{
	Do:          initial,
	Name:        "initial",
	Description: "Execute an initial deployment",
	Permission: "deploy#initial",
	MinArgs:    2,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

var DeployCmd cmd.Option = cmd.Option{
	Do:          deploy,
	Name:        "deploy",
	Description: "Execute a deployment, regardless of initial or upgrade",
	Permission: "deploy#initial",
	MinArgs:    2,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func initial(args intf.Command) (result string, err error) {
	return execute(args, "INITIAL")
}

func deploy(args intf.Command) (result string, err error) {
	return execute(args, "*")
}
