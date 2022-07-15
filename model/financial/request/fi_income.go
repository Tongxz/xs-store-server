package request

import (
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/financial"
)

type IncomeSearch struct {
	financial.Income
	request.PageInfo
}
