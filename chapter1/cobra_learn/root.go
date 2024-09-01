package cobra_learn

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Date string
)

var RootCmd = &cobra.Command{
	Use:   "mygo",
	Short: "我的GoClient",
	Long:  `我的go命令行客户端，用于模拟go命令行客户端`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	Version: "1.0",
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&Date, "date", "", "date")
	RootCmd.AddCommand(envCmd)
	RootCmd.AddCommand(getCmd)
	RootCmd.AddCommand(productCmd)
	RootCmd.AddCommand(clearCmd)
	RootCmd.AddCommand(consumeCmd)
	RootCmd.AddCommand(checkInfoCmd)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
