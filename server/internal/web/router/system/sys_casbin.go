package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewCasbinRouter(casbinApi *system.CasbinApi) *CasbinRouter {
	return &CasbinRouter{
		casbinApi: casbinApi,
	}
}

type CasbinRouter struct {
	casbinApi *system.CasbinApi
}

func (r *CasbinRouter) InitCasbinRouter(router *gin.RouterGroup, recordService *service.OperationRecordService) {
	casbinRouter := router.Group("casbin").Use(middleware.OperationRecord(recordService))
	casbinRouterWithoutRecord := router.Group("casbin")
	{
		casbinRouter.POST("updateCasbin", r.casbinApi.UpdateCasbin)
	}
	{
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", r.casbinApi.GetPolicyPathByAuthorityId)
	}
}
