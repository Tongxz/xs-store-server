package request

import (
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
