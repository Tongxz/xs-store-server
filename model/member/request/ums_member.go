package request

import (
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/member"
)

type MemberSearch struct {
	member.Member
	request.PageInfo
}
