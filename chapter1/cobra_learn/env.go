package cobra_learn

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "show env",
	Long:  "show all config data",
	Run: func(cmd *cobra.Command, args []string) {
		for _, item := range viper.AllKeys() {
			cmd.Println(viper.Get(item))
		}
	},
}
