package request

import (
	"github.com/tongxz/xs-admin-vue/server/model/common/request"
	"github.com/tongxz/xs-admin-vue/server/model/inventoryManage"
)

type WarehousingSearch struct {
	inventoryManage.Warehousing
	request.PageInfo
}
