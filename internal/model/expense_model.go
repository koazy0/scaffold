package model

import (
	"fmt"
)

type ExpenseModel struct {
	MODEL
	//Id           uint64              `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID       int64                `gorm:"column:user_id;not null;index"` // 外键字段
	User         *UserNoPasswordModel `gorm:"foreignKey:UserID;references:ID;"`
	Expense      Decimal              `gorm:"column:expense;type:decimal(10,2);size:64;comment:开支花费" json:"expense"`
	ExpenseName  string               `gorm:"column:config_key;size:64;comment:开支项" json:"expense_name"`
	ExpenseCycle int                  `gorm:"column:config_key;size:64;comment:开支周期" json:"expense_cycle"` //直接将花销除以周期
}

// cycle的周期
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

type UpdateExpenseReq struct {
	Request []struct {
		Id           uint64  `json:"id"`            //花销ID
		Expense      float64 `json:"expense"`       //花销
		ExpenseName  string  `json:"expense_name"`  //花销项
		ExpenseCycle int     `json:"expense_cycle"` //花销周期
		Option       string  `json:"option"`        //"add" 或者 "delete"
	} `json:"request"`
}

type UpdateExpenseReply struct {
	Message string `json:"message"`
}

type GetExpenseReq struct {
}

type GetExpenseReply struct {
	Expenses []struct {
		Id           uint64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
		Expense      Decimal `gorm:"column:expense;type:decimal(10,2);size:64;comment:开支花费" json:"expense"`
		ExpenseName  string  `gorm:"column:config_key;size:64;comment:开支项" json:"expense_name"`
		ExpenseCycle int     `gorm:"column:config_key;size:64;comment:开支周期" json:"expense_cycle"` //直接将花销除以周期
	} `json:"expenses"`
}
