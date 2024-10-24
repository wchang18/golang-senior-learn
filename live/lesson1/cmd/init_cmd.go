package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long:  "Initialize a new project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s init project success", projectName)
	},
}

var projectName string

func init() {
	RootCmd.AddCommand(InitCmd)
	InitCmd.Flags().StringVarP(&projectName, "name", "n", "", "project name")
}
