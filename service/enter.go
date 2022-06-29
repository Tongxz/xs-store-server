package service

import (
	"github.com/tongxz/xs-admin-vue/server/service/example"
	"github.com/tongxz/xs-admin-vue/server/service/financial"
	"github.com/tongxz/xs-admin-vue/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	FinancialServiceGroup financial.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
