package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/gin-gonic/gin"
)

func NewBaseRouter(baseApi *system.BaseApi) *BaseRouter {
	return &BaseRouter{
		baseApi: baseApi,
	}
}

type BaseRouter struct {
	baseApi *system.BaseApi
}

func (r *BaseRouter) InitBaseRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := router.Group("base")
	{
		baseRouter.POST("login", r.baseApi.Login)
		baseRouter.POST("captcha", r.baseApi.Captcha)
	}
	return baseRouter
}
