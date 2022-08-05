package request

import (
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/member"
)

type RechargeSearch struct {
	member.Recharge
	request.PageInfo
}
