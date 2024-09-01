package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	serv := g.Server()
	serv.BindHandler("GET:/order/{order_id}", GetOrder)
	serv.SetPort(8009)
	serv.Run()
}

func GetOrder(req *ghttp.Request) {
	req.Response.Writef("order:%v", req.Get("order_id"))
}
