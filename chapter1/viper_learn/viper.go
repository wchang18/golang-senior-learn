package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.StringP("name", "n", "default name", "your name")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	name := viper.Get("name")
	fmt.Println(name)
}
