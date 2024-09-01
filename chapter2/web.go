package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	httpServer := g.Server()
	httpServer.BindHandler("/", func(req *ghttp.Request) {
		req.Response.Write("Hello World!")
	})
	httpServer.SetPort(8000)
	httpServer.Run()
}
