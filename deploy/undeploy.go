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
	Description: "Uninstalls an application",
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
	// TODO Read error

	body, err = http.Create("/deployment", bytes.NewBuffer(body))
	// TODO Read error

	taskId := string(body)

	body, err = http.Create("/task/"+string(body)+"/start", nil)
	// TODO Read error

	displayStatus(taskId)

	return string(body), err
}
