package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"time"
)

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

func OrderMiddleware(r *ghttp.Request) {
	r.Response.Writeln("order middleware")
	r.Middleware.Next()
}

func main() {
	server := g.Server()
	server.Use(GlobalMiddleware)
	server.Group("api", func(group *ghttp.RouterGroup) {
		group.Middleware(UserMiddleware)
		group.GET("user/{id}", func(r *ghttp.Request) {
			time.Sleep(time.Second)
			r.Response.Writef("user id: %s", r.Get("id"))
		})

		group.Group("/order", func(group *ghttp.RouterGroup) {
			group.Middleware(OrderMiddleware)
			group.GET("/{id}", func(r *ghttp.Request) {
				time.Sleep(time.Millisecond * 10)
				r.Response.Writef("order id: %s", r.Get("id"))
			})
		})
	})
	server.SetPort(8009)
	server.Run()
}
