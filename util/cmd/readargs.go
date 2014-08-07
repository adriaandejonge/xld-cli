package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/adriaandejonge/xld/util/intf"
	"io/ioutil"
	"os"
	"strings"
)

func ReadArgs(args []string) intf.Command {
	indices := indexDashes(args)

	main := mainArgument(args)

	subs := subs(args, indices)

	var arguments = make([]intf.Argument, 0)

	if len(subs) > 0 && subs[len(subs)-1] == "stdin" {

		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			// TODO THROW ERROR
		}

		input := make(map[string]interface{})

		err = json.Unmarshal(bytes, &input)
		if err != nil {
			// TODO THROW ERROR
		}

		for k, v := range input {
			arguments = append(arguments, &JsonArg{k, v})
		}

	} else {
		arguments = cmdArguments(args, indices)
	}

	return &MainCmd{main, subs, arguments}
}

func indexDashes(args []string) (indices []int) {
	indices = make([]int, 0)
	for i, el := range args {
		if el[0] == 45 {
			indices = append(indices, i)
		}
	}
	return
}

func mainArgument(args []string) string {
	if len(args) > 0 {
		return args[0]
	} else {
		return ""
	}
}

func cmdArguments(args []string, indices []int) (arguments []intf.Argument) {
	arguments = make([]intf.Argument, len(indices))

	for i, index := range indices {

		var cmdArg *CmdArg

		if i == len(indices)-1 {
			cmdArg = &CmdArg{args[index][1:], args[index+1:]}
		} else {
			cmdArg = &CmdArg{args[index][1:], args[index+1 : indices[i+1]]}
		}

		// MOVE BELOW?
		arguments[i] = cmdArg

		if len(cmdArg.values) == 1 {

			value := cmdArg.values[0]
			input := strings.Split(value, ":")

			if len(input) == 2 && input[0] == "stdin" {

				switch input[1] {
				case "json":
					bytes, err := ioutil.ReadAll(os.Stdin)
					if err != nil {
						// TODO THROW ERROR
					}

					input := make(map[string]interface{})

					err = json.Unmarshal(bytes, &input)
					if err != nil {
						// TODO THROW ERROR
					}

					arguments[i] = &JsonArg{cmdArg.name, input}

				case "csv":
					fmt.Println("csv not yet implemented")

				}

			}

		}

	}
	return
}

func subs(args []string, indices []int) (subs []string) {
	if len(indices) > 0 {
		subs = args[1:indices[0]]
	} else if len(args) > 1 {
		subs = args[1:]
	}
	return
}
