package security

import (
 	 "encoding/xml"
	"fmt"
	"strings"

	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
)

var PermissionsCmd cmd.Option = cmd.Option{
	Do:          permissions,
	Name:        "permissions",
	Description: "Display available permissions",

	Permission: "",
	MinArgs:    0,
	Help: `
# xld permissions

Show a list of permissions that can be set using xld grant / xld revoke

Usage:

 - xld permissions
`,
}

func permissions(args intf.Command) (result string, err error) {
	body, err := http.Read("/metadata/permissions")
	if err != nil {
		return
	}

	list := PermissionsList{}
	err = xml.Unmarshal(body, &list)

	for _, el := range list.Permissions {
		fmt.Println(eq(el.Name, 20), eq(el.Level, 8), el.Root)
	}
		
	return "", nil
}

func eq(input string, desired int) string {
	repeat := desired - len(input)
	if repeat > 0 {
		return input + strings.Repeat(" ", repeat)
	}
	return input
}