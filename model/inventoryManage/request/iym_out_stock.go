package request

import (
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/inventoryManage"
)

type OutStockSearch struct {
	inventoryManage.OutStock
	request.PageInfo
}
