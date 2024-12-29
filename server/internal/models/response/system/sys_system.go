package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/config"
)

type SysConfigResponse struct {
	Config models.Server `json:"config"`
}
