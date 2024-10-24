package main

import "golang-senior-learn/live/lesson1/cmd"

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
