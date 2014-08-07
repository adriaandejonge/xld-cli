package deploy

import (
	"github.com/adriaandejonge/xld/util/cmd"
)

var (
	PlanCmd cmd.Option = cmd.Option{
		Do:          plan,
		Name:        "plan",
		Description: "Display steps in a deployment",
		Help: `
TODO: 
	Long, multi-line help text
`,
		Permission: "deploy#initial", // TODO Depends on initial / upgrade / remove
		MinArgs:    2,
	}

	DeployCmd cmd.Option = cmd.Option{
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

	UpgradeCmd cmd.Option = cmd.Option{
		Do:          upgrade,
		Name:        "upgrade",
		Description: "Updates an application deployment",
		Help: `
TODO: 
	Long, multi-line help text
`,
		Permission: "deploy#upgrade",
		MinArgs:    2,
	}

	UndeployCmd cmd.Option = cmd.Option{
		Do:          undeploy,
		Name:        "undeploy",
		Description: "Uninstalls an application",
		Help: `
TODO: 
	Long, multi-line help text
`,
		Permission: "deploy#undeploy",
		MinArgs:    1,
	}
)



