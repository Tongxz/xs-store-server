package request

import (
	"github.com/tongxz/xs-admin-vue/model/{{.Package}}"
	"github.com/tongxz/xs-admin-vue/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
