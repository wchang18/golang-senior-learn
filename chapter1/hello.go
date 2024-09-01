package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "my_app",
	Short: "我的app",
	Long:  `我的app，用于测试`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "hello",
	Long:  `hello`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")
	},
}

func init() {
	RootCmd.AddCommand(helloCmd)
}

func main() {
	RootCmd.Execute()
}
