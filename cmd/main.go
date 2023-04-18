package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"tdp-aiart/cmd/args"
	"tdp-aiart/cmd/initd"
	"tdp-aiart/cmd/subset"
)

var rcmd = &cobra.Command{
	Use:     "tdp-aiart",
	Short:   args.AppName,
	Long:    args.ReadmeText,
	Version: args.Version,
}

func init() {

	// 延迟执行

	cobra.OnInitialize(
		initd.Viper, initd.Dataset, initd.Logger,
	)

	// 全局参数

	rcmd.PersistentFlags().StringVarP(&initd.ViperFile, "config", "c", "config.yml", "config file path")

}

func Execute() {

	rcmd.AddCommand(
		subset.WithServer(), subset.WithUpdate(),
	)

	if err := rcmd.Execute(); err != nil {
		os.Exit(1)
	}

}
