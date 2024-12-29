package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request"
)

type JwtBlacklist struct {
	request.CommonModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
