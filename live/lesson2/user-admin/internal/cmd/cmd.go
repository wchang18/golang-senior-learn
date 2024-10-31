package cmd

import (
	"context"
	"user-admin/internal/controller/profile"
	"user-admin/internal/controller/user"
	"user-admin/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	_ "user-admin/internal/logic"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/api", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Session)
					group.Bind(
						user.NewV1(),
					)

					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(service.Middleware().Auth)
						group.Bind(profile.NewV1())
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
