// 自动生成模板Income
package financial

import (
	"github.com/tongxz/xs-admin-vue/global"
)

// Income 结构体
// 如果含有time.Time 请自行import time包
type Income struct {
	global.GVA_MODEL
	IncomeData    string         `json:"incomeData" form:"incomeData" gorm:"column:incomeData;comment:收入日期;"`
	Name          string         `json:"name" form:"name" gorm:"column:name;comment:姓名;"`
	Mobile        *int           `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号码;size:11;"`
	Amount        *float64       `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`
	Department    string         `json:"department" form:"department" gorm:"column:department;comment:部门;"`
	Category      *int           `json:"category" form:"category" gorm:"column:category;comment:类别;"`
	Payment       *int           `json:"payment" form:"payment" gorm:"column:payment;comment:收款方式;"`
	Invoice       *bool          `json:"invoice" form:"invoice" gorm:"column:invoice;comment:是否开票;"`
	Bill          string         `json:"bill" form:"bill" gorm:"column:bill;comment:发票号;"`
	Waiter        string         `json:"waiter" form:"waiter" gorm:"column:waiter;comment:负责人;"`
	Note          string         `json:"note" form:"note" gorm:"column:note;comment:备注;"`
	IncomeDetails []IncomeDetail `json:"incomeDetails"`
}
type IncomeDetail struct {
	global.GVA_MODEL
	IncomeID uint
	Type     string   `json:"type" gorm:"comment:收入类型;"`
	Sold     string   `json:"sold" gorm:"comment:售出物品;"`
	Number   *int     `json:"number" gorm:"comment:售出数量;"`
	Price    *float64 `json:"price" gorm:"comment:售出金额;"`
}

// TableName Income 表名
func (Income) TableName() string {
	return "income"
}
