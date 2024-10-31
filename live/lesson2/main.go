package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	server := g.Server()
	server.BindHandler("Get:/hello", func(req *ghttp.Request) {
		req.Response.Writeln("hello world")
	})
	server.SetPort(8001)
	server.Run()
}
