package metadata

import (
	"github.com/adriaandejonge/xld/util/cmd"
)

var (
	DescribeCmd cmd.Option = cmd.Option{
		Do:          describe,
		Name:        "describe",
		Description: "Describe properties for configuration type",
		Help: `
TODO: 
	Long, multi-line help text
`,
		Permission: "",
		MinArgs:    1,
	}

	TypesCmd cmd.Option = cmd.Option{
		Do:          types,
		Name:        "types",
		Description: "List configuration types",
		Help: `
TODO: 
	Long, multi-line help text
`,
		Permission: "",
		MinArgs:    0,
	}

)