package deploy

import (
	"encoding/xml"
	"fmt"

	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
)

var PlanDeployCmd cmd.Option = cmd.Option{
	Do:          planDeploy,
	Name:        "plan-deploy",
	Description: "Display steps in an initial deployment",

	Permission: "deploy#initial",
	MinArgs:    2,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

var PlanUpgradeCmd cmd.Option = cmd.Option{
	Do:          planUpgrade,
	Name:        "plan-upgrade",
	Description: "Display steps in an upgrade deployment",

	Permission: "deploy#update",
	MinArgs:    2,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func planDeploy(args intf.Command) (result string, err error) {
	return plan(args, "INITIAL")
}

func planUpgrade(args intf.Command) (result string, err error) {
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
