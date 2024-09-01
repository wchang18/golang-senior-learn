package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main2() {

	server := g.Server()
	server.BindHandler("/", func(req *ghttp.Request) {
		req.Response.Writefln("id:%v, name:%v", req.Get("id"), req.Get("name"))
	})

	server.SetPort(8008)
	server.Run()
}
