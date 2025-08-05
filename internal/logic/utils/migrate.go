package utils

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moyu/internal/common"
	"moyu/internal/model"
	"moyu/internal/service"
	"strings"
)

func init() {
	service.RegisterMigrations(Migrations())
}

// 负责迁移数据表等操作
type sMigrations struct {
}

var insMigrations = sMigrations{}

func Migrations() *sMigrations {
	return &insMigrations
}
func (s *sMigrations) Migrate(ctx context.Context) {

	dsnVar, err := g.Cfg().Get(ctx, "database.default.link")
	if err != nil {
		common.Logs().Fatal("no dsn configs: " + err.Error())
	}
	dsn := dsnVar.String()
	dsn = strings.TrimPrefix(dsn, "mysql:")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		common.Logs().Fatal("database connect error: " + err.Error())
	}

	// 自动迁移
	err = db.AutoMigrate(
		&model.UserModel{}, // 添加你所有要迁移的表结构
	)
	if err != nil {
		common.Logs().Fatal(err.Error())
	}

}
