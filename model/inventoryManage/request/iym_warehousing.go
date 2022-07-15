package request

import (
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/inventoryManage"
)

type WarehousingSearch struct {
	inventoryManage.Warehousing
	request.PageInfo
}
