package cmd

type (

	CsvArg struct {
		name string
		values [][]string
	}
)




func (a *CsvArg) Name() string {
	return a.name
}

func (a *CsvArg) Value() string {
	return "CSV IS ONLY FOR MAPs"
}

func (a *CsvArg) Values() []string {
	return []string{"CSV IS ONLY FOR MAPs"}
}

func (a *CsvArg) Map() map[string]string {
	
	newMap := make(map[string]string)

	for _, el := range a.values {
		if len(el) == 2 {
			newMap[el[0]] = el[1]
		}
	}

	return newMap
}
