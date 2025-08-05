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
	Id        uint64              `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID    uint                `gorm:"column:user_id;not null;index"` // 外键字段
	User      UserNoPasswordModel `gorm:"foreignKey:user_id;references:ID;constraint:OnDelete:SET NULL"`
	Key       string              `gorm:"column:config_key,size:64;comment:配置项 Key" json:"key"`
	Value     string              `gorm:"column:config_value,size:255;comment:配置项 Value"`
	CreatedAt time.Time           `gorm:"column:created_at,auto;comment:创建时间" json:"-" structs:"-"`
	DeletedAt time.Time           `gorm:"column:deleted_at,auto;comment:删除时间" json:"-" structs:"-"`
}
