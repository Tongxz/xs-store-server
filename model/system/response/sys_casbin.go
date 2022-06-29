package response

import (
	"github.com/tongxz/xs-admin-vue/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
