package model

import "time"

type MODEL struct {
	ID        int64     `gorm:"primaryKey;comment:id" json:"id,select($any)" structs:"-"`                       // 主键ID
	CreatedAt time.Time `gorm:"column:created_at;auto;comment:创建时间" json:"created_at,select($any)" structs:"-"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;auto;comment:更新时间" json:"-" structs:"-"`                       // 更新时间
	DeletedAt time.Time `gorm:"column:deleted_at;comment:删除时间" json:"-" structs:"-"`
}
