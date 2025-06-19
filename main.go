package main

import (
	"flag"
	_ "scaffold/internal/packed"
	"scaffold/internal/service"

	_ "scaffold/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"scaffold/internal/cmd"
)

func main() {
	var migrateFlag bool

	ctx := gctx.GetInitCtx()
	// 定义命令行参数
	// 这里后面放在一个专门的包里面
	flag.BoolVar(&migrateFlag, "migrate", false, "执行数据库迁移")
	flag.BoolVar(&migrateFlag, "m", false, "执行数据库迁移（简写）")
	flag.Parse()

	if migrateFlag {
		service.Migrations().Migrate(ctx)
	}
	cmd.Main.Run(ctx)
}
