package migrate

import "fmt"

func Run(ctx context.Context) error {
	dsn := "root:password@tcp(127.0.0.1:3306)/yourdb?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("数据库连接失败: %v", err)
	}

	// 自动迁移
	return db.AutoMigrate(
		&model.User{}, // 添加你所有要迁移的表结构
	)
}
