package main

import (
	"fmt"
	"github.com/adriaandejonge/xld/login"
	repo "github.com/adriaandejonge/xld/repository"
	"github.com/adriaandejonge/xld/metadata"
	"os"

	"github.com/adriaandejonge/xld/cmd"
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

	case "test":



	default:
		fmt.Println("XL Deploy Command Line Alternative - EXPERIMENTAL v0.1")
		fmt.Println("Created by Adriaan de Jonge - July, 2014")

	}
	if err != nil {
		fmt.Println("ERROR", err)	
	} else {
		fmt.Println("SUCCESS", result)
	}

}
