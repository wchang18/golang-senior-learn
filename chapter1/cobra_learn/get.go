package cobra_learn

import "github.com/spf13/cobra"

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get package",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("download " + args[0] + " success")
	},
}
