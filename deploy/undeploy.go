package deploy

import (
	"bytes"
	"github.com/adriaandejonge/xld/repo"
	"github.com/adriaandejonge/xld/util/cmd"
	"github.com/adriaandejonge/xld/util/http"
	"github.com/adriaandejonge/xld/util/intf"
)

var UndeployCmd cmd.Option = cmd.Option{
	Do:          undeploy,
	Name:        "undeploy",
	Description: "Uninstall an application",
	Permission: "deploy#undeploy",
	MinArgs:    1,
	Help: `
TODO: 
	Long, multi-line help text
`,
}

func undeploy(args intf.Command) (result string, err error) {
	subs := args.Subs()
	appToUndeploy := repo.AntiAbbreviate(subs[0])
	body, err := http.Read("/deployment/prepare/undeploy?deployedApplication=" + appToUndeploy)
	if err != nil {
		return
	}

	body, err = http.Create("/deployment", bytes.NewBuffer(body))
	if err != nil {
		return
	}

	taskId := string(body)

	body, err = http.Create("/task/"+string(body)+"/start", nil)
	if err != nil {
		return
	}

	displayStatus(taskId)

	return string(body), err
}
