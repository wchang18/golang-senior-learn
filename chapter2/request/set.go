package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main4() {
	s := g.Server()
	s.Group("/set", func(group *ghttp.RouterGroup) {
		group.Middleware(func(req *ghttp.Request) {
			req.SetParam("uid", 123)
			req.SetCtxVar("uname", "chang")
			req.Middleware.Next()
		})
		group.ALL("/param", func(req *ghttp.Request) {
			req.Response.Write(req.Get("uid"), req.GetCtxVar("uname"))
		})
	})
	s.SetPort(8008)
	s.Run()
}
