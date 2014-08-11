package main

import (
	"errors"
	"fmt"
	"github.com/adriaandejonge/xld/deploy"
	"github.com/adriaandejonge/xld/metadata"
	"github.com/adriaandejonge/xld/repo"
	"os"

	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/login"
)

var commands cmd.OptionList = cmd.OptionList{
	login.LoginCmd,
	//login.LoutoutCmd
	deploy.PlanCmd,
	deploy.DeployCmd,
	deploy.UpgradeCmd,
	deploy.UndeployCmd,
	repo.CreateCmd,
	repo.ReadCmd,
	repo.UpdateCmd,
	repo.RemoveCmd,
	repo.ListCmd,
	metadata.DescribeCmd,
	metadata.TypesCmd,
}

func main() {

	args, err := cmd.ReadArgs(os.Args[1:])
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)

	}

	var result string

	finder := commands.Finder()
	command, ok := finder(args.Main())

	if ok {

		if len(args.Subs()) >= command.MinArgs {
			result, err = command.Do(args)
		} else {
			errorText := fmt.Sprintf(
				"Command %s expects at least %d arguments",
				command.Name, command.MinArgs)
			err = errors.New(errorText)
		}
	} else {

		// TODO check update vs upgrade = similar
		// TODO Make list depend on permissions
		// TODO if not logged in or env var; show how

		fmt.Println("XL Deploy Command Line Alternative - EXPERIMENTAL v0.1")
		fmt.Println("Created by Adriaan de Jonge - August, 2014")

		fmt.Println("\nUsage: xld <command> <params...>\n\nCommands\n")

		for _, el := range commands.List() {

			name := el.Name + "            "
			name = name[:10]

			if !el.Hidden {
				fmt.Println(name, "-", el.Description)
			}
		}

		fmt.Println("\nFor additional help on parameters, type: xld help <command>")

	}
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	} else if result != "" {
		fmt.Println(result)
		os.Exit(0)
	}

}
