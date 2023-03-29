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
	Short:   "TDP Aiart",
	Long:    args.ReadmeText,
	Version: args.Version,
}

func init() {

	// 延迟执行

	cobra.OnInitialize(
		initd.Viper, initd.Dataset, initd.Logger,
	)

	// 全局参数

	rcmd.PersistentFlags().StringVarP(&initd.ViperFile, "config", "c", "./config.yml", "配置文件路径")

}

func Execute() {

	rcmd.AddCommand(
		subset.WithServer(),
	)

	if err := rcmd.Execute(); err != nil {
		os.Exit(1)
	}

}
