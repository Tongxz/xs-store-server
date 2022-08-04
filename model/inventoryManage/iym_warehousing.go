// 自动生成模板Warehousing
package inventoryManage

import (
	"github.com/tongxz/xs-admin-vue/global"
)

// Warehousing 结构体
// 如果含有time.Time 请自行import time包
type Warehousing struct {
	global.GVA_MODEL
	ImgUrl     string   `json:"imgUrl" form:"imgUrl" gorm:"column:img_url;comment:物品图片;"`
	Name       string   `json:"name" form:"name" gorm:"column:name;comment:入库物品名称;"`
	Department string   `json:"department" form:"department" gorm:"column:department;comment:所属部门;"`
	Type       *int     `json:"type" form:"type" gorm:"column:type;comment:物品所属分类;"`
	IncomeType *int     `json:"income_type" form:"income_type" gorm:"column:income_type;comment:物品所属分类;"`
	Payment    *int     `json:"payment" form:"payment" gorm:"column:payment;comment:支付方式;"`
	Quantity   *int     `json:"quantity" form:"quantity" gorm:"column:quantity;comment:入库物品数量;"`
	Margin     *int     `json:"margin" form:"margin" gorm:"column:margin;comment:剩余物品数量;"`
	Unit       string   `json:"unit" form:"unit" gorm:"column:unit;comment:物品单位;"`
	UnitPrice  *float64 `json:"unitPrice" form:"unitPrice" gorm:"column:unit_price;comment:物品单价;"`
	Cost       *float64 `json:"cost" form:"cost" gorm:"column:cost;comment:成本价;"`
	Amount     *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:总金额;"`
	Remarks    string   `json:"remarks" form:"remarks" gorm:"column:remarks;comment:入库备注/说明;"`
}
type WarehousingName struct {
	global.GVA_MODEL
	Name   string `json:"name" form:"name" gorm:"column:name;comment:入库物品名称;"`
	Margin *int   `json:"margin" form:"margin" gorm:"column:margin;comment:剩余物品数量;"`
	Unit   string `json:"unit" form:"unit" gorm:"column:unit;comment:剩余物品数量;"`
}
type WarehousingUp struct {
	global.GVA_MODEL
	ImgUrl     []string `json:"imgUrl" `
	Name       string   `json:"name" `
	Department string   `json:"department" `
	Type       *int     `json:"type" `
	Payment    *int     `json:"payment" `
	Quantity   *int     `json:"quantity" `
	Margin     *int     `json:"margin" `
	Unit       string   `json:"unit" `
	UnitPrice  *float64 `json:"unitPrice" `
	Amount     *float64 `json:"amount" `
	Remarks    string   `json:"remarks" `
}
type WarehousingExcel struct {
	FileName string        `json:"fileName"` // 文件名
	InfoList []Warehousing `json:"infoList"`
}

// TableName Warehousing 表名
func (Warehousing) TableName() string {
	return "iym_Warehousing"
}

// TableName Warehousing 表名
func (WarehousingName) TableName() string {
	return "iym_Warehousing"
}
