package cmd

import (
	"context"
	"scaffold/internal/controller/ping"
	"scaffold/internal/controller/user"
	"scaffold/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					//service.Middleware().Ctx, //todo 在这个ctx中存储一系列消息
					service.Middleware().MiddlewareHandlerResponse,
				)
				group.Bind(
					ping.NewV1().Ping,
					user.NewV1().SignUp,
					user.NewV1().SignIn,
				)
				group.Group("user", func(groupUser *ghttp.RouterGroup) {
					groupUser.Middleware(service.Middleware().AccessKeyAuth)
					group.Bind()
				})
			})
			s.Run()
			return nil
		},
	}
)
