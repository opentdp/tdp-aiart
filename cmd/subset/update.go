package subset

import (
	"github.com/opentdp/go-helper/logman"
	"github.com/opentdp/go-helper/upgrade"
	"github.com/spf13/cobra"

	"tdp-aiart/cmd/args"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update assistant",
	Long:  "TDP Aiart Update Assistant",
	Run: func(cmd *cobra.Command, rq []string) {
		ExecUpdate()
	},
}

func WithUpdate() *cobra.Command {

	return updateCmd

}

func ExecUpdate() error {

	err := upgrade.Apply(&upgrade.RequesParam{
		Server:  args.UpdateUrl,
		Version: args.Version,
	})

	if err == nil {
		logman.Info("Update Success")
	}

	return err

}
