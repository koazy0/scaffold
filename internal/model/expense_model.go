package model

import (
	"fmt"
	"time"
)

type ExpenseModel struct {
	MODEL
	Id           uint64              `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID       uint                `gorm:"column:user_id;not null;index"` // 外键字段
	User         UserNoPasswordModel `gorm:"foreignKey:user_id;references:ID;constraint:OnDelete:SET NULL"`
	Expense      Decimal             `gorm:"column:expense;type:decimal(10,2);size:64;comment:开支花费" json:"expense"`
	ExpenseName  string              `gorm:"column:config_key;size:64;comment:开支项" json:"expense_name"`
	ExpenseCycle int                 `gorm:"column:config_key;size:64;comment:开支项" json:"expense_cycle"`
	CreatedAt    time.Time           `gorm:"column:created_at;auto;comment:创建时间" json:"-" structs:"-"`
	DeletedAt    time.Time           `gorm:"column:deleted_at;auto;comment:删除时间" json:"-" structs:"-"`
}

const (
	Daily = 1 << iota
	Weekly
	Monthly
)

type Decimal float64

// MarshalJSON 输出一个无引号的数字文本,保留两位小数
// e.g.:  "123.45"
func (d Decimal) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%.2f", float64(d))
	return []byte(s), nil
}
