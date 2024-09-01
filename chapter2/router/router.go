package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	serv := g.Server()
	serv.BindHandler("GET:/:name", func(req *ghttp.Request) {
		req.Response.Writef("url:%s,name:%s", req.URL.Path, req.Get("name"))
	})
	serv.BindHandler("POST:/list/{page}.html", func(req *ghttp.Request) {
		req.Response.Writef("url:%s,page:%s", req.URL.Path, req.Get("page"))
	})
	serv.BindHandler("PUT:/category/*any", func(req *ghttp.Request) {
		req.Response.Writef("url:%s,any:%s", req.URL.Path, req.Get("any"))
	})

	serv.SetPort(8009)
	serv.Run()
}
