package main

import (
	"fmt"
	"github.com/adriaandejonge/xld/login"
	repo "github.com/adriaandejonge/xld/repository"
	"github.com/adriaandejonge/xld/metadata"
	"os"
)

func main() {
	fmt.Println("XL Deploy Command Line Alternative - EXPERIMENTAL v0.1")
	fmt.Println("Created by Adriaan de Jonge - July, 2014")
	var err error
	var result string

	switch os.Args[1] {
	case "login", "connect":
		result, err = login.Do(os.Args[2:])
	case "create", "update", "remove", "list":
		result, err = repo.Do(os.Args[1:])
	case "types", "describe":
		result, err = metadata.Do(os.Args[1:])
	default:
		fmt.Println("Command not recognized")


	}
	if err != nil {
		fmt.Println("ERROR", err)	
	} else {
		fmt.Println("SUCCESS", result)
	}

}
