package request

import (
	"github.com/tongxz/xs-admin-vue/server/model/common/request"
	"github.com/tongxz/xs-admin-vue/server/model/inventoryManage"
)

type OutStockSearch struct {
	inventoryManage.OutStock
	request.PageInfo
}
