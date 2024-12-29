package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/config"
)

// 配置文件结构体
type System struct {
	Config models.Server `json:"config"`
}
