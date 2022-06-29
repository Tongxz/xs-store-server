package response

import "github.com/tongxz/xs-admin-vue/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
