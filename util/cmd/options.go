package cmd

import (
	"fmt"
	"github.com/adriaandejonge/xld/util/intf"
)

type (
	Executer func(args intf.Command) (result string, err error)
	Finder   func(command string) (option Option, ok bool)

	OptionList []Option

	Option struct {
		Do          Executer
		Name        string
		Description string
		Help        string
		Permission  string // TODO []string instead?
		MinArgs     int
		Hidden      bool
	}
)

func (optionList *OptionList) Finder() Finder {

	index := make(map[string]Option)

	var HelpCmd Option = Option{
		Do: func() Executer {
			return func(args intf.Command) (result string, err error) {
				subs := args.Subs()
				if len(subs) > 0 {
					if subs[0] == "all" {
						for _, el := range index {
							fmt.Println(el.Help)
						}
					}

					// TODO: Loop over multiple
					// TODO: Loop over all
					option, ok := index[subs[0]]
					if ok {
						fmt.Println(option.Help)
					}
				}
				return
			}
		}(),
		Name:        "help",
		Description: "Additional help for commands",
		Permission:  "",
		MinArgs:     1,
		Hidden:      true,
	}
	optionList.add(&HelpCmd)

	for _, el := range *optionList {
		index[el.Name] = el
	}

	return func(command string) (option Option, ok bool) {
		val, ok := index[command]
		return val, ok
	}
}

func (optionList *OptionList) List() (options []Option) {
	options = make([]Option, 0)
	for _, el := range *optionList {
		// TODO if permission ok
		options = append(options, el)
	}
	return
}

func (optionList *OptionList) add(option *Option) {
	*optionList = append(*optionList, *option)
}
