package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/guid"
)

func main() {
	serv := g.Server()
	serv.BindHandler("GET:/order/{order_id}", GetOrder)
	serv.BindHandler("POST:/order", AddOrder)
	serv.SetPort(8009)
	serv.Run()
}

func GetOrder(req *ghttp.Request) {
	req.Response.Writef("order:%v", req.Get("order_id"))
}

func AddOrder(req *ghttp.Request) {
	req.Response.Writef("add order, id:%v", guid.S())
}
