package main

import (
	_ "github.com/go-sql-driver/mysql"              // 必须匿名导入！  database/sql 驱动
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // GoFrame ORM adapter
	_ "moyu/internal/packed"

	"moyu/internal/cmd"
	_ "moyu/internal/logic"
)

func main() {
	cmd.Execute()
}
