// 自动生成模板Expenses
package financial

import (
	"github.com/tongxz/xs-admin-vue/global"
	"time"
)

// Expenses 结构体
// 如果含有time.Time 请自行import time包
type Expenses struct {
	global.GVA_MODEL
	ExpnDate   *time.Time `json:"expnDate" form:"expnDate" gorm:"column:expnDate;comment:收入日期;"`
	Content    string     `json:"content" form:"content" gorm:"column:content;comment:支出内容说明;"`
	Amount     *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:支出金额;"`
	Department string     `json:"department" form:"department" gorm:"column:department;comment:部门;"`
	Type       *int       `json:"type" form:"type" gorm:"column:type;comment:支出类型;"`
	Payment    *int       `json:"payment" form:"payment" gorm:"column:payment;comment:支付方式;"`
	Executor   string     `json:"executor" form:"executor" gorm:"column:executor;comment:执行该操作的人员;"`
	Invoice    *bool      `json:"invoice" form:"invoice" gorm:"column:invoice;comment:是否开票;"`
	Note       string     `json:"note" form:"note" gorm:"column:note;comment:备注说明;"`
}

// TableName Expenses 表名
func (Expenses) TableName() string {
	return "expenses"
}
