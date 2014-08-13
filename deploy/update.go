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
# xld update <app id> <env id>

Executes an update deployment. This explicitly does *not* work for initial deployments. Use xld deploy to deploy regardless of intitial/upgrade.

Usage:

 - xld update <app id> <env id>

Examples:

 - xld update app/MyApp/2.0 env/MyEnv
`,
}

func update(args intf.Command) (result string, err error) {
	return execute(args, "UPDATE")
}
