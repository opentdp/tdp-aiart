package service

import (
	"os"

	"github.com/kardianos/service"
	"github.com/opentdp/go-helper/logman"

	"tdp-aiart/service/server"
)

var statusMap = map[service.Status]string{
	0: "Unknown",
	1: "Running",
	2: "Stopped",
}

func Control(name, act string) {

	var svc service.Service

	// 获取抽象类

	switch name {
	case "server":
		svc = server.Service(cliArgs())
	default:
		logman.Fatal("unknown service", "name", name)
	}

	// 执行服务动作

	switch act {
	case "":
		if err := svc.Run(); err != nil {
			logman.Fatal(svc.String(), "run", "failed", "error", err)
		}
	case "status": // 查看状态
		if sta, err := svc.Status(); err != nil {
			logman.Fatal(svc.String(), act, "failed", "error", err)
		} else {
			logman.Warn(svc.String(), "status", statusMap[sta])
		}
	default: // 其他动作
		if err := service.Control(svc, act); err != nil {
			logman.Fatal(svc.String(), act, "failed", "error", err)
		}
	}

}

func cliArgs() []string {

	args := []string{}

	for i, n := 1, len(os.Args); i < n; i++ {
		if v := os.Args[i]; v != "-s" && v != "--service" {
			args = append(args, v)
		} else {
			i++
		}
	}

	return args

}
