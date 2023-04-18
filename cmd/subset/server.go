package subset

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tdp-aiart/cmd/args"
	"tdp-aiart/service"
)

var serverAct string

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Manage the server",
	Long:  "TDP Aiart Server Management",
	Run: func(cmd *cobra.Command, rq []string) {
		args.SubCommand.Name = cmd.Name()
		args.SubCommand.Action = serverAct
		service.Control(args.SubCommand.Name, serverAct)
	},
}

func WithServer() *cobra.Command {

	serverCmd.Flags().StringVarP(&serverAct, "service", "s", "", "management server service")
	serverCmd.Flags().StringP("listen", "l", ":7700", "server listens to ip addresse and port")

	viper.BindPFlag("server.listen", serverCmd.Flags().Lookup("listen"))

	return serverCmd

}
