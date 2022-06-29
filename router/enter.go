package router

import (
	"github.com/tongxz/xs-admin-vue/server/router/example"
	"github.com/tongxz/xs-admin-vue/server/router/financial"
	"github.com/tongxz/xs-admin-vue/server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Financial financial.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
