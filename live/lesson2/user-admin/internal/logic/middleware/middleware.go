package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gutil"
	"net/http"
	"user-admin/internal/consts"
	"user-admin/internal/service"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

func (m *sMiddleware) CORS(req *ghttp.Request) {
	req.Response.CORSDefault()
	req.Middleware.Next()
}

func (m *sMiddleware) Session(req *ghttp.Request) {
	service.Session().Init(req)
	req.Middleware.Next()
}

func (m *sMiddleware) Auth(req *ghttp.Request) {
	username := service.Session().Get(consts.UserNameKey)
	fmt.Println("t1:", username)
	if gutil.IsEmpty(username) {
		req.Response.WriteStatus(http.StatusForbidden)
		return
	}
	req.SetCtxVar(consts.UserNameKey, username)
	req.Middleware.Next()
}
