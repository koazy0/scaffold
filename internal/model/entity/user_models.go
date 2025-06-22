// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserModels is the golang structure for table user_models.
type UserModels struct {
	Id        uint64      `json:"id"        orm:"id"         ` // id
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	Uid       string      `json:"uid"       orm:"uid"        ` // UID
	UserId    string      `json:"userId"    orm:"user_id"    ` // 用户ID
	Username  string      `json:"username"  orm:"username"   ` // 用户姓名
	Salt      string      `json:"salt"      orm:"salt"       ` // 密码盐值
	Password  string      `json:"password"  orm:"password"   ` // 密码
	Role      int         `json:"role"      orm:"role"       ` // 权限，1管理员，2普通用户，3游客
	Status    int64       `json:"status"    orm:"status"     ` // 注册来源，1qq，3邮箱
	DeleteAt  *gtime.Time `json:"deleteAt"  orm:"delete_at"  ` // 删除时间
}
