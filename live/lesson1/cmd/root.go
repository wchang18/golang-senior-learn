package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "gosys",
	Short: "gosys is a system information tool",
	Long:  `gosys is a system information tool`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	Version: "1.0.0",
}
