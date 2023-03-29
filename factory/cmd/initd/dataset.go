package initd

import (
	"os"

	"tdp-aiart/cmd/args"
)

func Dataset() {

	datadir := args.Dataset.Dir

	if datadir != "" && datadir != "." {
		os.MkdirAll(datadir, 0755)
	}

}
