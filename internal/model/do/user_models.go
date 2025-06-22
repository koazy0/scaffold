// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserModels is the golang structure of table user_models for DAO operations like Where/Data.
type UserModels struct {
	g.Meta    `orm:"table:user_models, do:true"`
	Id        interface{} // id
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	Uid       interface{} // UID
	UserId    interface{} // 用户ID
	Username  interface{} // 用户姓名
	Salt      interface{} // 密码盐值
	Password  interface{} // 密码
	Role      interface{} // 权限，1管理员，2普通用户，3游客
	Status    interface{} // 注册来源，1qq，3邮箱
	DeleteAt  *gtime.Time // 删除时间
}
