package request

import (
	"github.com/tongxz/xs-admin-vue/server/model/common/request"
	"github.com/tongxz/xs-admin-vue/server/model/financial"
)

type ExpensesSearch struct {
	financial.Expenses
	request.PageInfo
}
