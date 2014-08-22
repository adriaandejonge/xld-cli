package cmd

import (
	"encoding/json"
	"github.com/adriaandejonge/xld/util/intf"
)

type (
	StdinCmd struct {
		main      string
		values    map[string]interface{}
	}

	JsonArg struct {
		name string
		arg  interface{}
	}
)
/*
func NewStdinCmd(main string, values map[string]interface{}) *StdinCmd {
	return &StdinCmd{main, values}
}*/

func NewStdinCmd(main string, jsonStr string) (stdinCmd *StdinCmd, err error) {
	if err != nil {
		return nil, err
	}

	values := make(map[string]interface{})

	err = json.Unmarshal([]byte(jsonStr), &values)
	if err != nil {
		return nil, err
	}
	
	
	return &StdinCmd{main, values}, nil
}

func (s *StdinCmd) Main() string {
	return s.main
}

func (s *StdinCmd) Subs() []string {
	// TODO Do you want this knowledge here???
	switch s.main {
	case "create", "modify-sub":
		return []string{
			s.values["type"].(string),
			s.values["id"].(string),
		}
	case "modify":
		return []string{
			s.values["id"].(string),
		}
	default:
		// TODO Throw error: not supported
		return make([]string, 0)
	}

}

func (s *StdinCmd) Arguments() []intf.Argument {
	var arguments = make([]intf.Argument, 0)

	for k, v := range s.values["content"].(map[string]interface{}) {
		arguments = append(arguments, &JsonArg{k, v})
	}
	return arguments
}

func (a *JsonArg) Name() string {
	return a.name
}

func (a *JsonArg) Value() string {
	if str, ok := a.arg.(string); ok {
		return str
	} else {
		return "NOT A STRING VALUE"
	}

}

func (a *JsonArg) Values() []string {
	if intfArr, ok := a.arg.([]interface{}); ok {
		strArr := make([]string, len(intfArr))
		for i, el := range intfArr {
			strArr[i] = el.(string)
		}
		return strArr
	} else {
		return make([]string, 0)
		// TODO or throw error?
	}
}

func (a *JsonArg) Map() map[string]string {
	if strMap, ok := a.arg.(map[string]interface{}); ok {
		newMap := make(map[string]string)

		for k, v := range strMap {
			newMap[k] = v.(string)
		}
		return newMap
	} else {
		return make(map[string]string)
		// TODO or throw error?
	}
}
