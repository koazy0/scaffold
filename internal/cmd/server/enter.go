package server

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/spf13/cobra"
	"scaffold/internal/common"
	"scaffold/internal/controller/ping"
	"scaffold/internal/controller/user"
	"scaffold/internal/service"
)

func ServerCommand() *cobra.Command {
	return serverCmd
}

var logger = common.Logs().Cat("cmd/server")

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  `just run server as normal`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		ctx := gctx.GetInitCtx()
		logger.Infoln("start server")
		Main.Run(ctx)
	},
}

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
					ping.NewV1(),
				)
				group.Group("/user", func(groupUser *ghttp.RouterGroup) {
					//groupUser.Middleware(service.Middleware().AccessKeyAuth)
					groupUser.Bind(
						user.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
