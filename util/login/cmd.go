package login

import (
	"github.com/adriaandejonge/xld/util/cmd"
)

var (
	LoginCmd cmd.Option = cmd.Option{
		Do:          Do,
		Name:        "login",
		Description: "Provide URL, username and password",
		Permission: "",
		MinArgs:    3,
		Help: `
# XLD Login: 

Provide login details for XL Deploy server. Credentials are stored base64 encoded in a .xld file in the root of your user profile for reuse in subsequent requests.

Usage:

 - xld login <server> <username> <password>

Example:

 - xld login localhost:4516 admin $ecr3tP@ssw0rd
`,
	}
)