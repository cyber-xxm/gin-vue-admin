package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/gin-gonic/gin"
)

func NewJwtRouter(jwtApi *system.JwtApi) *JwtRouter {
	return &JwtRouter{
		jwtApi: jwtApi,
	}
}

type JwtRouter struct {
	jwtApi *system.JwtApi
}

func (r *JwtRouter) InitJwtRouter(router *gin.RouterGroup) {
	jwtRouter := router.Group("jwt")
	{
		jwtRouter.POST("jsonInBlacklist", r.jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
