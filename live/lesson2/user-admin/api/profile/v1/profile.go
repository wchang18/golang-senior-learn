package v1

import "github.com/gogf/gf/v2/frame/g"

type GetProfileReq struct {
	g.Meta `path:"/profile" tags:"Profile" method:"get" summary:"Get Profile"`
}

type GetProfileRes struct {
	MyInfo
}

type MyInfo struct {
	Username  string `json:"username"`
	Telephone string `json:"telephone"`
}
