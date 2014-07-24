package intf

type (

	Command interface {

		Main() string
		Subs() []string
		Arguments() []Argument

	}
	Argument interface {

		Name() string
		Value() string
		Values() []string
		Map() map[string]string

	}
)