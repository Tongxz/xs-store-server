package service

import (
	"github.com/tongxz/xs-admin-vue/server/service/example"
	"github.com/tongxz/xs-admin-vue/server/service/financial"
	"github.com/tongxz/xs-admin-vue/server/service/inventoryManage"
	"github.com/tongxz/xs-admin-vue/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup          system.ServiceGroup
	ExampleServiceGroup         example.ServiceGroup
	FinancialServiceGroup       financial.ServiceGroup
	InventorymanageServiceGroup inventoryManage.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
