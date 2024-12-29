package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
)

type SysAPIResponse struct {
	Api system.SysApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []system.SysApi `json:"apis"`
}

type SysSyncApis struct {
	NewApis    []system.SysApi `json:"newApis"`
	DeleteApis []system.SysApi `json:"deleteApis"`
}
