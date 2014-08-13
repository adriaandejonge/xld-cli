package deploy

import (
	"encoding/xml"
	"fmt"

	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
)

var PlanInitialCmd cmd.Option = cmd.Option{
	Do:          planInitial,
	Name:        "plan-initial",
	Description: "Display steps in an initial deployment",

	Permission: "deploy#initial",
	MinArgs:    2,
	Help: `
# XLD Plan-initial: 

Show the steps for an initial deployment without executing. For execution, see xld initial

Usage:

 - xld plan-initial <app id> <env id>

Example:

 - xld plan-initial app/MyApp/1.0 env/MyEnv
`,
}

var PlanUpdateCmd cmd.Option = cmd.Option{
	Do:          planUpdate,
	Name:        "plan-update",
	Description: "Display steps in an update deployment",

	Permission: "deploy#update",
	MinArgs:    2,
	Help: `
# XLD Plan-update: 

Show the steps for an update deployment without executing. For execution, see xld update

Usage:

 - xld plan-update <app id> <env id>

Example:

 - xld plan-update app/MyApp/2.0 env/MyEnv
`,
}

func planInitial(args intf.Command) (result string, err error) {
	return plan(args, "INITIAL")
}

func planUpdate(args intf.Command) (result string, err error) {
	return plan(args, "UPDATE")
}


func plan(args intf.Command, depType string) (result string, err error) {
	result, err = prepare(args, depType)

	body, err := http.Read("/task/" + result + "/step")
	if err != nil {
		return
	}

	task := Task{}
	err = xml.Unmarshal(body, &task)
	if err != nil {
		return "error", err
	}

	fmt.Println("Plan", task.Description)

	for i, step := range task.Steps {
		fmt.Printf("%d/%d - "+step.Description+"\n", i+1, task.TotalSteps)
	}

	return "", err
}
