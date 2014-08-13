package metadata

import (
	"fmt"
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/intf"
)

var DescribeCmd cmd.Option = cmd.Option{
	Do:          describe,
	Name:        "describe",
	Description: "Describe properties for configuration type",
	Permission: "",
	MinArgs:    1,
	Help: `
# XLD Describe: 

Print properties and property type for item type(s).

Usage:

 - xld describe <item type(s)

Examples:

 - xld describe jee.War

 - xld describe tomcat.Server udm.Directory

 - xld describe $(xld types | grep tomcat)
`,
}

func describe(args intf.Command) (result string, err error) {
	subs := args.Subs()

	for _, sub := range subs {
		_, err := describeOne(sub)
		if err != nil {
			return "error", err
		}
	}
	return "", nil
}

func describeOne(typeName string) (result string, err error) {
	ciType, err := Type(typeName)
	if err != nil {
		return "error", err
	}

	fmt.Println(ciType.Type + ":")

	for _, prop := range ciType.Properties {
		fmt.Println("  -"+prop.Name, prop.Kind, iif(prop.Required, "required", ""), iif(prop.Hidden, "hidden", ""))
	}

	return "", nil
}
