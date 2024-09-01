package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UserController struct {
}

func (c *UserController) SignUp(req *ghttp.Request) {
	req.Response.Write("注册成功")
}

func (c *UserController) SignIn(req *ghttp.Request) {
	req.Response.Write("登录成功")
}

func main() {
	serv := g.Server()
	//serv.SetNameToUriType(ghttp.UriTypeCamel)
	//serv.BindObject("/user", new(UserController))
	serv.BindObjectMethod("/user-sign-in", new(UserController), "SignIn")
	serv.SetPort(8009)
	serv.Run()
}
