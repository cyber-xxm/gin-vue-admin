package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request/system"
)

type PolicyPathResponse struct {
	Paths []system.CasbinInfo `json:"paths"`
}
