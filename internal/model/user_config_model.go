package model

import (
	"time"
)

// UserNoPasswordModel 主表，存一个用户名
type UserNoPasswordModel struct {
	//g.Meta `orm:"table:user_nopassword"`

	ID        uint      `json:"id,select($any)" gorm:"primaryKey;comment:id" structs:"-"`
	UserID    *string   `json:"user_id" gorm:"column:user_id;type:varchar(64);uniqueIndex;not null;comment:用户ID"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at;comment:新增时间"  structs:"-"`
	DeletedAt time.Time `json:"-" gorm:"column:deleted_at;comment:删除时间"  structs:"-"`

	// 反向关联，当查询user的时候直接一次性拉取所有的配置和花销
	Configs []UserConfigModel `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// UserConfigModel 配置表模型，子表，通过 user_id 关联到 user.id
type UserConfigModel struct {
	MODEL
	Id            int64               `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID        int64               `gorm:"column:user_id;not null;index"` // 外键字段
	User          UserNoPasswordModel `gorm:"foreignKey:user_id;references:ID;constraint:OnDelete:SET NULL"`
	Income        Decimal             `gorm:"column:income;comment:薪资" json:"income"`
	IncomeCycle   int                 `gorm:"column:income_cycle;comment:上班周期" json:"income_cycle"`               //默认0~30
	WorkTimeStart string              `gorm:"column:work_time_start,size:64;comment:上班时间" json:"work_time_start"` //只取time.TimeOnly 格式
	WorkTimeEnd   string              `gorm:"column:work_time_end,size:255;comment:下班时间 " json:"work_time_end"`
	CreatedAt     time.Time           `gorm:"column:created_at,auto;comment:创建时间" json:"-" structs:"-"`
	DeletedAt     time.Time           `gorm:"column:deleted_at,auto;comment:删除时间" json:"-" structs:"-"`
}

type GetConfigReq struct {
}

type GetConfigReply struct {
	Expenses struct {
		Id           uint64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		Expense      Decimal `gorm:"column:expense;type:decimal(10,2);size:64;comment:开支花费" json:"expense"`
		ExpenseName  string  `gorm:"column:config_key;size:64;comment:开支项" json:"expense_name"`
		ExpenseCycle int     `gorm:"column:config_key;size:64;comment:开支周期" json:"expense_cycle"` //直接将花销除以周期
	} `json:"config"`
}

type UpdateConfigReq struct {
	Request struct {
		Id           uint64  `json:"id"`            //花销ID
		Expense      float64 `json:"expense"`       //花销
		ExpenseName  string  `json:"expense_name"`  //花销项
		ExpenseCycle int     `json:"expense_cycle"` //花销周期
		Option       string  `json:"option"`        //"add" 或者 "delete"
	} `json:"request"`
}

type UpdateConfigReply struct {
}
