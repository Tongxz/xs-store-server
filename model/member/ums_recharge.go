// 自动生成模板Recharge
package member

import (
	"github.com/tongxz/xs-admin-vue/global"
)

// Recharge 结构体
// 如果含有time.Time 请自行import time包
type Recharge struct {
	global.GVA_MODEL
	MemberId *int     `json:"memberId" form:"memberId" gorm:"column:member_id;comment:会员ID;"`
	Amount   *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:充值金额;"`
	Sign     *bool    `json:"sign" form:"sign" gorm:"column:sign;comment:充值金额;"`
}

// TableName Recharge 表名
func (Recharge) TableName() string {
	return "ums_recharge"
}
