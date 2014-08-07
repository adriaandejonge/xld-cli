package cmd

type (
	JsonArg struct {
		name string
		arg  interface{}
	}
)

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
	if strArr, ok := a.arg.([]string); ok {
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
