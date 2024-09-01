package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func main3() {

	server := g.Server()
	server.Use(ghttp.MiddlewareHandlerResponse)
	server.Group("/user", func(group *ghttp.RouterGroup) {
		group.Bind(new(controller))
	})
	server.SetPort(8008)
	server.Run()

}

type SignUpReq struct {
	g.Meta   `path:"/sign-up" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	UserName string `v:"required|email#用户必填|用户名为邮箱"`
	Pass     string `p:"password" v:"required|length:6,16"`
	Pass2    string `p:"password2" v:"required|length:6,16|same:Pass"`
	NickName string `d:"default name"`
}

type SignUpRes struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
}

type controller struct {
}

func (c *controller) SignUp(ctx context.Context, req *SignUpReq) (res *SignUpRes, err error) {

	glog.Info(ctx, "password:", req.Pass, req.Pass2)
	res = &SignUpRes{
		Id:       1,
		UserName: req.UserName,
		NickName: req.NickName,
	}
	return
}
