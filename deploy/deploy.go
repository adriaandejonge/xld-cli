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
# XLD Initial: 

Executes an initial deployment. This explicitly does *not* work for upgrade deployments. Use xld deploy to deploy regardless of intitial/upgrade.

Usage:

 - xld initial <app id> <env id>

 - xld initial <app id> <env id> -orchestrator <orchestrator(s)>

Examples:

 - xld initial app/MyApp/1.0 env/MyEnv

 - xld initial app/MyComp/1.0 env/MyEnv -orchestrator parallel-by-container parallel-by-composite-package
`,
}

var DeployCmd cmd.Option = cmd.Option{
	Do:          deploy,
	Name:        "deploy",
	Description: "Execute a deployment, regardless of initial or upgrade",
	Permission: "deploy#initial",
	MinArgs:    2,
	Help: `
# XLD Deploy: 

Executes a deployment, either initial or update. If you need to make a distinction, use xld initial or xld update instead.

Usage:

 - xld deploy <app id> <env id>

 - xld deploy <app id> <env id> -orchestrator <orchestrator(s)>

Examples:

 - xld deploy app/MyApp/2.0 env/MyEnv

 - xld deploy app/MyComp/3.0 env/MyEnv -orchestrator parallel-by-container parallel-by-composite-package
`,
}

func initial(args intf.Command) (result string, err error) {
	return execute(args, "INITIAL")
}

func deploy(args intf.Command) (result string, err error) {
	return execute(args, "*")
}
