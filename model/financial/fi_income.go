// 自动生成模板Income
package financial

import (
	"github.com/tongxz/xs-admin-vue/server/global"
)

// Income 结构体
// 如果含有time.Time 请自行import time包
type Income struct {
	global.GVA_MODEL
	Name     string   `json:"name" form:"name" gorm:"column:name;comment:姓名;"`
	Mobile   *int     `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号码;size:11;"`
	Amount   *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`
	Category string   `json:"category" form:"category" gorm:"column:category;comment:类别;"`
	Payment  string   `json:"payment" form:"payment" gorm:"column:payment;comment:收款方式;"`
	Invoice  *bool    `json:"invoice" form:"invoice" gorm:"column:invoice;comment:是否开票;"`
	Waiter   string   `json:"waiter" form:"waiter" gorm:"column:waiter;comment:负责人;"`
	Note     string   `json:"note" form:"note" gorm:"column:note;comment:备注;"`
}

// TableName Income 表名
func (Income) TableName() string {
	return "income"
}
