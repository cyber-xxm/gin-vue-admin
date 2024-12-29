package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request"
	"time"
)

type SysParamsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Name           string     `json:"name" form:"name" `
	Key            string     `json:"key" form:"key" `
	request.PageInfo
}
