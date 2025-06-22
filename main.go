package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"              // 必须匿名导入！  database/sql 驱动
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // GoFrame ORM adapter
	"github.com/gogf/gf/v2/frame/g"
	"scaffold/internal/dao"
	"scaffold/internal/model/do"
	_ "scaffold/internal/packed"
	"scaffold/internal/service"

	_ "scaffold/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"scaffold/internal/cmd"
)

func main() {
	var migrateFlag bool

	ctx := gctx.GetInitCtx()
	va, err := g.Cfg().Get(ctx, "database")
	fmt.Println(va.String())
	va, err = g.Cfg().Get(ctx, "database.default.link")
	fmt.Println(va.String())
	//sql.Register("mysql", &mysql.MySQLDriver{})
	// 测试数据库连接
	if db := g.DB(); db != nil {
		if _, err := db.Exec(ctx, "SELECT 1"); err != nil {
			g.Log().Fatalf(ctx, "数据库连接测试失败:%v", err)
		}
	}
	// 定义命令行参数
	// 这里后面放在一个专门的包里面
	flag.BoolVar(&migrateFlag, "migrate", false, "执行数据库迁移")
	flag.BoolVar(&migrateFlag, "m", false, "执行数据库迁移（简写）")
	flag.Parse()

	user := do.UserModels{}
	users := []do.UserModels{}
	err = dao.UserModels.Ctx(ctx).Where("id=1").Scan(&user)
	count := 1
	err = dao.UserModels.Ctx(ctx).Where("id=1").ScanAndCount(&users, &count, false)

	if err != nil {
		fmt.Println(err.Error())
	}

	if migrateFlag {
		service.Migrations().Migrate(ctx)
	}
	cmd.Main.Run(ctx)
}
