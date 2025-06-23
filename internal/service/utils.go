// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IExample interface {
		Example()
	}
	IJwt interface {
		// GenerateToken 生成一个带有 userID的 JWT，默认有效期 12 小时
		GenerateToken(ctx context.Context, userID string) (string, error)
		ParseToken(tokenString string) (UserID string, err error)
		// HashPassword 用 SHA-256(salt + password) 方式加密
		HashPassword(password string, salt string) (passwordEncrypt string)
		// GenerateSalt 生成随机盐值，长度固定为16个字节
		GenerateSalt() (string, error)
	}
	IMigrations interface {
		Migrate(ctx context.Context)
	}
)

var (
	localExample    IExample
	localJwt        IJwt
	localMigrations IMigrations
)

func Example() IExample {
	if localExample == nil {
		panic("implement not found for interface IExample, forgot register?")
	}
	return localExample
}

func RegisterExample(i IExample) {
	localExample = i
}

func Jwt() IJwt {
	if localJwt == nil {
		panic("implement not found for interface IJwt, forgot register?")
	}
	return localJwt
}

func RegisterJwt(i IJwt) {
	localJwt = i
}

func Migrations() IMigrations {
	if localMigrations == nil {
		panic("implement not found for interface IMigrations, forgot register?")
	}
	return localMigrations
}

func RegisterMigrations(i IMigrations) {
	localMigrations = i
}
