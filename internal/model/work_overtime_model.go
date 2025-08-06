package model

import "time"

type WorkOverTimeModel struct {
	MODEL
	Id             uint64              `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID         uint                `gorm:"column:user_id;not null;index"` // 外键字段
	User           UserNoPasswordModel `gorm:"foreignKey:user_id;references:ID;constraint:OnDelete:SET NULL"`
	OverTimeIncome Decimal             `gorm:"column:expense;type:decimal(10,2);size:64;comment:加班报酬" json:"over_time_income"` //可为正可为负，前后进行
	Ratio          Decimal             `gorm:"column:ratio;type:decimal(10,2);size:64;comment:加班报酬倍率" json:"ratio"`            //最多两位小数
	HasIncome      bool                `gorm:"column:has_income;size:64;comment:是否有额外报酬" json:"has_income"`                    //最多两位小数
	StartTime      time.Time           `gorm:"column:start_time;auto;comment:加班开始时间" json:"start_time"`
	EndTime        time.Time           `gorm:"column:end_time;auto;comment:加班结束时间" json:"end_time"`
	CreatedAt      time.Time           `gorm:"column:created_at;auto;comment:创建时间" json:"datetime" structs:"-"` //这里只需要date
	DeletedAt      time.Time           `gorm:"column:deleted_at;auto;comment:删除时间" json:"-" structs:"-"`
}

type UpdateWorkOverTimeReq struct {
	Request struct {
		Id             uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		OverTimeIncome Decimal   `gorm:"column:expense;type:decimal(10,2);size:64;comment:加班报酬" json:"over_time_income"` //可为正可为负，前后进行
		Ratio          Decimal   `gorm:"column:ratio;type:decimal(10,2);size:64;comment:加班报酬倍率" json:"ratio"`            //最多两位小数
		HasIncome      bool      `gorm:"column:has_income;size:64;comment:是否有额外报酬" json:"has_income"`                    //最多两位小数
		StartTime      time.Time `gorm:"column:start_time;auto;comment:加班开始时间" json:"start_time"`
		EndTime        time.Time `gorm:"column:end_time;auto;comment:加班结束时间" json:"end_time"`
	} `json:"request"`
}

type UpdateWorkOverTimeReply struct {
	Message string `json:"message"`
}

type GetWorkOverTimeReq struct {
}

type GetWorkOverTimeReply struct {
	WorkOverTime struct {
		Id             uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		OverTimeIncome Decimal   `gorm:"column:expense;type:decimal(10,2);size:64;comment:加班报酬" json:"over_time_income"` //可为正可为负，前后进行
		Ratio          Decimal   `gorm:"column:ratio;type:decimal(10,2);size:64;comment:加班报酬倍率" json:"ratio"`            //最多两位小数
		HasIncome      bool      `gorm:"column:has_income;size:64;comment:是否有额外报酬" json:"has_income"`                    //最多两位小数
		StartTime      time.Time `gorm:"column:start_time;auto;comment:加班开始时间" json:"start_time"`
		EndTime        time.Time `gorm:"column:end_time;auto;comment:加班结束时间" json:"end_time"`
	} `json:"work_over_time"`
}
