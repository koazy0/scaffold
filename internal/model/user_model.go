package model

import (
	"moyu/internal/model/ctype"
	"time"
)

// UserModel 数据库中用户表
// 这里暂时先弃用
type UserModel struct {
	MODEL
	UID      *string          `json:"uid" gorm:"column:uid;uniqueIndex;not null;type:varchar(36);comment:UID"`
	UserID   *string          `json:"user_id" gorm:"column:user_id;type:varchar(64);uniqueIndex;not null;comment:用户ID"`
	Username *string          `json:"username" gorm:"column:username;type:varchar(64);comment:用户姓名"`
	Salt     *string          `json:"salt" gorm:"column:salt;type:varchar(128);comment:密码盐值"`
	Password *string          `json:"password" gorm:"column:password;type:varchar(128);comment:密码"`
	Role     ctype.Role       `json:"type" gorm:"column:role;size:4;default:1;comment:权限，1管理员，2普通用户，3游客"`
	Status   ctype.SignStatus `json:"status" gorm:"column:status;type:int;default:1;comment:注册来源，1qq，3邮箱"`
	DeleteAt time.Time        `gorm:"comment:删除时间" json:"-" structs:"-"`
}

type UserSignIn struct {
	UserID string `json:"user_id" v:"required"`
}

type UserSignInReply struct {
	Token string `json:"token"`
}

type UserLogout struct {
	UserId string `json:"username" v:"required"`
}

type UserLogoutReply struct {
	Message string `json:"message"`
}
