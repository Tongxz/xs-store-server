package service

import (
	"github.com/tongxz/xs-admin-vue/service/example"
	"github.com/tongxz/xs-admin-vue/service/financial"
	"github.com/tongxz/xs-admin-vue/service/inventoryManage"
	"github.com/tongxz/xs-admin-vue/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup          system.ServiceGroup
	ExampleServiceGroup         example.ServiceGroup
	FinancialServiceGroup       financial.ServiceGroup
	InventorymanageServiceGroup inventoryManage.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
