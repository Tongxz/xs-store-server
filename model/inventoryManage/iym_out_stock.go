// 自动生成模板OutStock
package inventoryManage

import (
	"github.com/tongxz/xs-admin-vue/server/global"
)

// OutStock 结构体
// 如果含有time.Time 请自行import time包
type OutStock struct {
	global.GVA_MODEL
	Type       string   `json:"type" form:"type" gorm:"column:type;comment:出库类型;"`
	Name       string   `json:"name" form:"name" gorm:"column:name;comment:出库物品名称;"`
	Department string   `json:"department" form:"department" gorm:"column:department;comment:部门;"`
	Item_type  string   `json:"item_type" form:"item_type" gorm:"column:item_type;comment:物品类型;"`
	Unit       string   `json:"unit" form:"unit" gorm:"column:unit;comment:物品单位;"`
	Quantity   *int     `json:"quantity" form:"quantity" gorm:"column:quantity;comment:出库数量;"`
	Unit_price *float64 `json:"unit_price" form:"unit_price" gorm:"column:unit_price;comment:物品单价;"`
	Amount     *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:总金额;"`
	Remarks    string   `json:"remarks" form:"remarks" gorm:"column:remarks;comment:备注;"`
}

// TableName OutStock 表名
func (OutStock) TableName() string {
	return "iym_out_stock"
}
