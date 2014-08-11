package deploy

import (
	"encoding/xml"
	"fmt"

	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
)

var PlanCmd cmd.Option = cmd.Option{
	Do:          plan,
	Name:        "plan",
	Description: "Display steps in a deployment",

	Permission: "deploy#initial", // TODO Depends on initial / upgrade / remove
	MinArgs:    2,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func plan(args intf.Command) (result string, err error) {
	result, err = prepare(args)

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
