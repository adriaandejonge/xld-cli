package main

import (
	"fmt"
	"os"
	"github.com/adriaandejonge/xld/repo"
	"github.com/adriaandejonge/xld/metadata"
	"github.com/adriaandejonge/xld/deploy"
	
	"github.com/adriaandejonge/xld/util/login"
	"github.com/adriaandejonge/xld/util/cmd"
)

func main() {

	args := cmd.ReadArgs(os.Args[1:])


	var err error
	var result string

	switch args.Main() {
	case "login", "connect":
		result, err = login.Do(args)
	case "create", "update", "remove", "list":
		result, err = repo.Do(args)
	case "types", "describe":
		result, err = metadata.Do(args)
	case "plan", "deploy", "upgrade", "undeploy":
		result, err = deploy.Do(args)
	default:

		// TODO check update vs upgrade = similar
		
		fmt.Println("XL Deploy Command Line Alternative - EXPERIMENTAL v0.1")
		fmt.Println("Created by Adriaan de Jonge - July, 2014")
		fmt.Println("\nUsage: xld <command> <params...>\n\nCommands\n")
		fmt.Println("login    - Provide URL, username and password")
		fmt.Println("create   - Create new configuration item")
		fmt.Println("update   - Change existing configuration item")
		fmt.Println("remove   - Remove existing configuration item")
		fmt.Println("list     - List configuration items")
		fmt.Println("types    - List configuration types")
		fmt.Println("describe - Describe properties for configuration type")
		fmt.Println("plan     - Display steps in a deployment")
		fmt.Println("deploy   - Execute a deployment")
		fmt.Println("upgrade  - Updates an application deployment")
		fmt.Println("undeploy - Uninstalls an application")
		fmt.Println("\nFor additional help on parameters, type: xld <command> help")


	}
	if err != nil {
		fmt.Println("ERROR", err)	
	} else if result != "" {
		fmt.Println(result)
	}

}
