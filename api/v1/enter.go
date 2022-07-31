package v1

import (
	"github.com/tongxz/xs-admin-vue/api/v1/example"
	"github.com/tongxz/xs-admin-vue/api/v1/financial"
	"github.com/tongxz/xs-admin-vue/api/v1/inventoryManage"
	"github.com/tongxz/xs-admin-vue/api/v1/member"
	"github.com/tongxz/xs-admin-vue/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup          system.ApiGroup
	ExampleApiGroup         example.ApiGroup
	FinancialApiGroup       financial.ApiGroup
	InventorymanageApiGroup inventoryManage.ApiGroup
	MemberApiGroup          member.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
