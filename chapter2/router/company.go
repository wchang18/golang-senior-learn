package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type OrderReq struct {
	g.Meta `path:"/order" method:"get"`
	Id     int    `v:"required" dc:"order id"`
	No     string `v:"size:8"`
}

type OrderRes struct {
	Id   int    `json:"id"`
	No   string `json:"no"`
	Name string `json:"name"`
}

type Order struct {
}

func (c Order) GetOrder(ctx context.Context, req *OrderReq) (res *OrderRes, err error) {
	res = &OrderRes{
		Id:   req.Id,
		No:   req.No,
		Name: "order name",
	}
	return
}

func main() {
	s := g.Server()
	s.Use(ghttp.MiddlewareHandlerResponse)
	s.Group("api", func(group *ghttp.RouterGroup) {
		group.Bind(new(Order))
	})
	s.SetPort(8009)
	s.Run()
}
