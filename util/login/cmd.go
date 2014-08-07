package login

import (
	"github.com/adriaandejonge/xld/util/cmd"
)

var (
	LoginCmd cmd.Option = cmd.Option{
		Do:          Do,
		Name:        "login",
		Description: "Provide URL, username and password",
		Help: `
TODO: 
	Long, multi-line help text
`,
		Permission: "",
		MinArgs:    3,
	}
)