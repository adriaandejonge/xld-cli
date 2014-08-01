package cmd

import (
	"strings"
	"github.com/adriaandejonge/xld/util/intf"
)


func ReadArgs(args []string) intf.Command {
	indices := make([]int, 0)
	for i, el := range args {
		if el[0] == 45 {
			indices = append(indices, i)
		}
	}

	var main string

	if len(args) > 0 {
		main = args[0]
	} else {
		main = ""
	}

	var subs []string
	if len(indices) > 0 {
		subs = args[1:indices[0]]	
	} else if len(args) > 1 {
		subs = args[1:]
	}
	

	arguments := make([]intf.Argument, len(indices))

	for i, index := range indices {
		
		if i == len(indices) - 1 {
			arguments[i] = &Argument{args[index][1:], args[index+1:]}
		} else {
			arguments[i] = &Argument{args[index][1:], args[index+1:indices[i+1]]}
		}

	}

	return &MainCmd{main, subs, arguments}
}

type(

	MainCmd struct {
		main string
		subs []string
		arguments []intf.Argument
	}

	Argument struct {
		name string
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

func (a *Argument) Name() string {
	return a.name
}

func (a *Argument) Value() string {
	return a.values[0]
}

func (a *Argument) Values() []string {
	return a.values
}

func (a *Argument) Map() map[string]string {
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

		