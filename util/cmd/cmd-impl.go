package cmd

import (
	"github.com/adriaandejonge/xld/util/intf"
	"strings"
)

type (
	MainCmd struct {
		main      string
		subs      []string
		arguments []intf.Argument
	}

	CmdArg struct {
		name   string
		values []string
	}
)

func (m *MainCmd) Main() string {
	return m.main
}

func (m *MainCmd) Subs() []string {
	return m.subs
}

func (m *MainCmd) Arguments() []intf.Argument {
	return m.arguments
}

func (a *CmdArg) Name() string {
	return a.name
}

func (a *CmdArg) Value() string {
	return a.values[0]
}

func (a *CmdArg) Values() []string {
	return a.values
}

func (a *CmdArg) Map() map[string]string {
	argMap := make(map[string]string)
	for _, arg := range a.values {
		kv := strings.SplitN(arg, "=", 2)
		if len(kv) > 1 {
			argMap[kv[0]] = kv[1]
		} else {
			argMap[kv[0]] = ""
		}
	}
	return argMap
}
