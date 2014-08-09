package deploy

import (
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
)

var DeployCmd cmd.Option = cmd.Option{
	Do:          deploy,
	Name:        "deploy",
	Description: "Execute a deployment",
	Help: `
TODO: 
	Long, multi-line help text
`,
	Permission: "deploy#initial",
	MinArgs:    2,
}

func deploy(args intf.Command) (result string, err error) {
	result, err = prepare(args)

	body, err := http.Create("/task/"+result+"/start", nil)

	displayStatus(result)

	return string(body), err
}
