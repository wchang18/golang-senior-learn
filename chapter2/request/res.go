package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main5() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/json", func(r *ghttp.Request) {
			r.Response.WriteJson(g.Map{
				"name": "tom",
				"age":  10,
			})
		})
	})
	s.SetPort(8008)
	s.Run()
}

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/{name}", func(r *ghttp.Request) {
			if r.Get("name").String() == "tom" {
				r.Response.Write("name is tom")
				r.Exit()
			}
			r.Response.Writef("name : %v", r.Get("name"))
		})
	})
	s.SetPort(8008)
	s.Run()
}
