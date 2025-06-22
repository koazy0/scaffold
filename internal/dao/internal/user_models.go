// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserModelsDao is the data access object for the table user_models.
type UserModelsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserModelsColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserModelsColumns defines and stores column names for the table user_models.
type UserModelsColumns struct {
	Id        string // id
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	Uid       string // UID
	UserId    string // 用户ID
	Username  string // 用户姓名
	Salt      string // 密码盐值
	Password  string // 密码
	Role      string // 权限，1管理员，2普通用户，3游客
	Status    string // 注册来源，1qq，3邮箱
	DeleteAt  string // 删除时间
}

// userModelsColumns holds the columns for the table user_models.
var userModelsColumns = UserModelsColumns{
	Id:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Uid:       "uid",
	UserId:    "user_id",
	Username:  "username",
	Salt:      "salt",
	Password:  "password",
	Role:      "role",
	Status:    "status",
	DeleteAt:  "delete_at",
}

// NewUserModelsDao creates and returns a new DAO object for table data access.
func NewUserModelsDao(handlers ...gdb.ModelHandler) *UserModelsDao {
	return &UserModelsDao{
		group:    "default",
		table:    "user_models",
		columns:  userModelsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserModelsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserModelsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserModelsDao) Columns() UserModelsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserModelsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserModelsDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserModelsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
