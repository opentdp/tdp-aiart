package initd

import (
	"os"

	"tdp-aiart/cmd/args"
	"tdp-aiart/helper/logman"
)

func Logger() {

	logdir := args.Logger.Dir

	if logdir != "" && logdir != "." {
		os.MkdirAll(logdir, 0755)
	}

	logman.New()

}
