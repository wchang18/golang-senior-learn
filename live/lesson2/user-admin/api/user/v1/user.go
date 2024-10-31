package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" tags:"Login" method:"post" summary:"Login"`
	Username string `json:"username" v:"required|length:6,30#请输入账号|账号长度为6~30位"`
	Password string `json:"password" v:"required|password"`
}

type LoginRes struct{}

type RegisterReq struct {
	g.Meta    `path:"/register" tags:"Register" method:"post" summary:"Register"`
	Username  string `json:"username" v:"required|length:6,30#请输入账号|账号长度为6~30位"`
	Password  string `json:"password" v:"required|password"`
	Password2 string `json:"password2" v:"required|same:Password#请再次输入密码|两次密码输入不一致"`
	Telephone string `json:"telephone" v:"telephone#请输入正确手机号"`
}

type RegisterRes struct {
	Username  string `json:"username"`
	Telephone string `json:"telephone"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" tags:"Logout" method:"post" summary:"Logout"`
}

type LogoutRes struct{}
