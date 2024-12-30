package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewCasbinRouter(casbinApi *system.CasbinApi, recordService *service.OperationRecordService) *CasbinRouter {
	return &CasbinRouter{
		casbinApi: casbinApi,
	}
}

type CasbinRouter struct {
	casbinApi     *system.CasbinApi
	recordService *service.OperationRecordService
}

func (r *CasbinRouter) InitCasbinRouter(router *gin.RouterGroup) {
	casbinRouter := router.Group("casbin").Use(middleware.OperationRecord(r.recordService))
	casbinRouterWithoutRecord := router.Group("casbin")
	{
		casbinRouter.POST("updateCasbin", r.casbinApi.UpdateCasbin)
	}
	{
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", r.casbinApi.GetPolicyPathByAuthorityId)
	}
}
