package deploy

import (
	"errors"
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/intf"
)

var UpgradeCmd cmd.Option = cmd.Option{
	Do:          upgrade,
	Name:        "upgrade",
	Description: "Updates an application deployment",
	Permission: "deploy#upgrade",
	MinArgs:    2,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func upgrade(args intf.Command) (result string, err error) {
	return "error", errors.New("Update is not yet implemented")
}
