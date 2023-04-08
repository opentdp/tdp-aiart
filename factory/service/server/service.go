package server

import (
	"github.com/kardianos/service"

	"tdp-aiart/cmd/args"
	"tdp-aiart/helper/logman"
)

var svclog service.Logger

func Service(param []string) service.Service {

	config := &service.Config{
		Name:        "tdp-aiart",
		DisplayName: "TDP Aiart Server",
		Description: "TDP Control Panel Server",
		Option: service.KeyValue{
			"LogDirectory": args.Logger.Dir,
			"LogOutput":    args.Logger.Target == "file",
		},
		Arguments: param,
	}

	svc, err := service.New(&program{}, config)
	if err != nil {
		logman.Fatal("Init service failed", "error", err)
	}

	svclog, err = svc.Logger(nil)
	if err != nil {
		logman.Fatal("Init service failed", "error", err)
	}

	return svc

}
