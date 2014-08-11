package cmd

import (
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


func (s *StdinCmd) Main() string {
	return s.main
}

func (s *StdinCmd) Subs() []string {
	return []string{
		s.values["type"].(string),
		s.values["id"].(string),
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
