package deploy

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/adriaandejonge/xld/repo"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"

	"github.com/clbanning/mxj/j2x"

	"strings"
	"time"
)

func prepare(args intf.Command) (result string, err error) {
	subs := args.Subs()
	appVersion := repo.AntiAbbreviate(subs[0])
	targetEnv := repo.AntiAbbreviate(subs[1]) // or 2?

	parts := strings.Split(appVersion, "/")

	app := parts[len(parts)-2]
	//version := parts[len(parts) - 1]

	deployment := map[string]interface{}{
		"deployment": map[string]interface{}{
			"-type": "INITIAL",
			"application": map[string]interface{}{
				"udm.DeployedApplication": map[string]interface{}{
					"-id": targetEnv + "/" + app,
					"version": map[string]interface{}{
						"-ref": appVersion,
					},
					"environment": map[string]interface{}{
						"-ref": targetEnv,
					},
					"optimizePlan": "true",
				},
			},
		},
	}

	// TODO Make this a util?
	json, _ := j2x.MapToJson(deployment)
	xml, _ := j2x.JsonToXml(json)

	body, err := http.Create("/deployment/prepare/deployeds", bytes.NewBuffer(xml))
	if err != nil {
		return "error", err
	}

	body, err = http.Create("/deployment", bytes.NewBuffer(body))

	return string(body), err
}

func displayStatus(taskId string) {
	timer := time.Tick(100 * time.Millisecond)

	previousStep := -1

	for _ = range timer {

		// TODO support parallel deployments
		body, err := http.Read("/task/" + taskId + "/step")
		// TODO Read error

		task := Task{}
		err = xml.Unmarshal(body, &task)
		if err != nil {
			return
		}

		currentStep := task.CurrentStep

		for i, step := range task.Steps {

			// Fix current step - sometimes it is one too high
			if step.State == "EXECUTING" {
				currentStep = i
			}

			if i > previousStep && i <= currentStep {
				fmt.Printf("%d/%d - "+step.Description+"\n", i+1, task.TotalSteps)
			}
		}

		previousStep = currentStep

		if task.State == "EXECUTED" {
			break
		} else if task.State == "FAILED" {
			// TODO throw error on deployment FAILED
			// TODO show logs on deployment error
			// can we simulate an error?
			break
		}

	}
}
