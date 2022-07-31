// 自动生成模板Member
package member

import (
	"github.com/tongxz/xs-admin-vue/global"
	"time"
)

// Member 结构体
// 如果含有time.Time 请自行import time包
type Member struct {
	global.GVA_MODEL
	Name          string     `json:"name" form:"name" gorm:"column:name;comment:会员姓名;"`
	Mobile        *int       `json:"mobile" form:"mobile" gorm:"column:mobile;comment:会员手机号;"`
	Gender        *int       `json:"gender" form:"gender" gorm:"column:gender;comment:会员性别;"`
	Like          string     `json:"like" form:"like" gorm:"column:like;comment:会员喜好;"`
	Diet          string     `json:"diet" form:"diet" gorm:"column:diet;comment:会员喜好;"`
	MemberCard    *bool      `json:"memberCard" form:"memberCard" gorm:"column:member_card;comment:是否会员;"`
	MemberBalance *float64   `json:"memberBalance" form:"memberBalance" gorm:"column:member_balance;comment:会员卡余额;"`
	OpenDate      *time.Time `json:"openDate" form:"openDate" gorm:"column:open_date;comment:会员卡开通时间;"`
}

// TableName Member 表名
func (Member) TableName() string {
	return "ums_member"
}
