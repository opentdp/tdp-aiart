package main

import (
	"embed"

	"tdp-aiart/cmd"
	"tdp-aiart/cmd/args"
)

//go:embed front
var efs embed.FS

func main() {

	args.Efs = &efs

	cmd.Execute()

}
