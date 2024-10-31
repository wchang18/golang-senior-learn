package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"time"
)

func main() {
	s := g.Server()
	s.Use(GlobalMiddleware)
	s.Group("/api", func(group *ghttp.RouterGroup) {

		group.ALL("/all", func(r *ghttp.Request) {
			r.Response.Write("all")
		})
		group.GET("/get", func(r *ghttp.Request) {
			time.Sleep(time.Second)
			r.Response.Write("get")
		})
		group.POST("/post", func(r *ghttp.Request) {
			r.Response.Write("post")
		})
		
		s.Group("/api/v1", func(group *ghttp.RouterGroup) {
			group.Middleware(UserMiddleware)
			group.PUT("/put", func(r *ghttp.Request) {
				r.Response.Write("put")
			})
		})
	})
	s.SetPort(8009)
	s.Run()
}

func GlobalMiddleware(r *ghttp.Request) {
	glog.Info(context.Background(), "global middleware")
	start := time.Now().UnixMilli()
	r.Middleware.Next()
	end := time.Now().UnixMilli()
	glog.Info(context.Background(), "执行时间（毫秒）：", end-start)
}

func UserMiddleware(r *ghttp.Request) {
	r.Response.Writeln("user middleware")
	r.Middleware.Next()
}
