package request

import (
	"github.com/tongxz/xs-admin-vue/server/model/{{.Package}}"
	"github.com/tongxz/xs-admin-vue/server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
