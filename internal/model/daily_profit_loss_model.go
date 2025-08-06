package model

import "time"

// DailyProfitLossModel 每日营收汇总
type DailyProfitLossModel struct {
	MODEL
	Id        uint64              `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID    uint                `gorm:"column:user_id;not null;index"`
	User      UserNoPasswordModel `gorm:"foreignKey:user_id;references:ID;constraint:OnDelete:SET NULL"`
	Date      time.Time           `gorm:"column:date;type:date;not null;comment:日期" json:"date"`           //只需要每日计算,计算当日时间点
	Salary    Decimal             `gorm:"column:salary;type:decimal(10,2);comment:基础工资" json:"salary"`     //通过MySQL触发器进行自动更新，
	Overtime  Decimal             `gorm:"column:overtime;type:decimal(10,2);comment:加班收入" json:"overtime"` //通过MySQL触发器进行自动更新
	Expense   Decimal             `gorm:"column:expense;type:decimal(10,2);comment:总开销" json:"expense"`    //通过MySQL触发器进行自动更新，由ExpenseModel进行触发
	Profit    Decimal             `gorm:"column:profit;type:decimal(10,2);comment:盈亏" json:"profit"`       //通过MySQL触发器进行自动更新，由profit进行触发
	CreatedAt time.Time           `gorm:"column:created_at;auto;comment:创建时间" json:"-"`
	DeletedAt time.Time           `gorm:"column:deleted_at;auto;comment:删除时间" json:"-"`
}
