package request

import (
	"github.com/tongxz/xs-admin-vue/server/model/common/request"
	"github.com/tongxz/xs-admin-vue/server/model/financial"
)

type IncomeSearch struct {
	financial.Income
	request.PageInfo
}
