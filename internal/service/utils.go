// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IJwt interface {
		// GenerateToken 生成一个带有 userID 和 username 的 JWT，默认有效期 1 小时
		GenerateToken(userID string, password string) (string, error)
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
	localJwt        IJwt
	localMigrations IMigrations
)

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
