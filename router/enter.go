package router

import (
	"github.com/tongxz/xs-admin-vue/router/example"
	"github.com/tongxz/xs-admin-vue/router/financial"
	"github.com/tongxz/xs-admin-vue/router/inventoryManage"
	"github.com/tongxz/xs-admin-vue/router/member"
	"github.com/tongxz/xs-admin-vue/router/system"
)

type RouterGroup struct {
	System          system.RouterGroup
	Example         example.RouterGroup
	Financial       financial.RouterGroup
	Inventorymanage inventoryManage.RouterGroup
	Member          member.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
