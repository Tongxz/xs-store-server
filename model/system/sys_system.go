package system

import (
	"github.com/tongxz/xs-admin-vue/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
