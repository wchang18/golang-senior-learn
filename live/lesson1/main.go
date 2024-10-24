package main

import (
	"flag"
	"fmt"
)

func main2() {
	var name string
	flag.StringVar(&name, "name", "", "名称")

	flag.Parse()

	if name != "" {
		fmt.Println("name:", name)
	} else {
		fmt.Println("name is empty")
	}
}
