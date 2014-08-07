package repo

import (
	"github.com/adriaandejonge/xld/util/cmd"
)

var (

	CreateCmd cmd.Option = cmd.Option{
		Do:          create,
		Name:        "create",
		Description: "Create new configuration item",
		Help: `
TODO: 
	Long, multi-line help text
`,
		Permission: "repo#edit",
		MinArgs:    0,
	}

	UpdateCmd cmd.Option = cmd.Option{
		Do:          update,
		Name:        "update",
		Description: "Change existing configuration item",
		Help: `
TODO: 
	Long, multi-line help text
`,
		Permission: "repo#edit",
		MinArgs:    0,
	}

	RemoveCmd cmd.Option = cmd.Option{
		Do:          remove,
		Name:        "remove",
		Description: "Remove existing configuration item",
		Help: `
TODO: 
	Long, multi-line help text
`,
		Permission: "import#remove", //repo#edit
		MinArgs:    0,
	}

	ListCmd cmd.Option = cmd.Option{
		Do:          list,
		Name:        "list",
		Description: "List configuration items",
		Help: `
TODO: 
	Long, multi-line help text
`,
		Permission: "",
		MinArgs:    0,
	}

	// TODO ReadCmd permission: read
)