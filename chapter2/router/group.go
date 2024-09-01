package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	serv := g.Server()
	serv.Group("/api/v2", func(group *ghttp.RouterGroup) {
		group.GET("user/{id}", func(req *ghttp.Request) {
			req.Response.Writef("获取用户: %v", req.Get("id"))
		})
		group.POST("user", func(req *ghttp.Request) {
			req.Response.Write("添加用户信息")
		})
		group.ALL("user/profile", func(req *ghttp.Request) {
			req.Response.Write("用户信息")
		})
	})
	serv.SetPort(8009)
	serv.Run()
}
