package v1

import (
	"github.com/tongxz/xs-admin-vue/server/api/v1/example"
	"github.com/tongxz/xs-admin-vue/server/api/v1/financial"
	"github.com/tongxz/xs-admin-vue/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	FinancialApiGroup financial.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
