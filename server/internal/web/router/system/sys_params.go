package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewSysParamsRouter(sysParamsApi *system.SysParamsApi, recordService *service.OperationRecordService) *SysParamsRouter {
	return &SysParamsRouter{
		sysParamsApi:  sysParamsApi,
		recordService: recordService,
	}
}

type SysParamsRouter struct {
	sysParamsApi  *system.SysParamsApi
	recordService *service.OperationRecordService
}

// InitSysParamsRouter 初始化 参数 路由信息
func (r *SysParamsRouter) InitSysParamsRouter(router *gin.RouterGroup) {
	sysParamsRouter := router.Group("sysParams").Use(middleware.OperationRecord(r.recordService))
	sysParamsRouterWithoutRecord := router.Group("sysParams")
	{
		sysParamsRouter.POST("createSysParams", r.sysParamsApi.CreateSysParams)             // 新建参数
		sysParamsRouter.DELETE("deleteSysParams", r.sysParamsApi.DeleteSysParams)           // 删除参数
		sysParamsRouter.DELETE("deleteSysParamsByIds", r.sysParamsApi.DeleteSysParamsByIds) // 批量删除参数
		sysParamsRouter.PUT("updateSysParams", r.sysParamsApi.UpdateSysParams)              // 更新参数
	}
	{
		sysParamsRouterWithoutRecord.GET("findSysParams", r.sysParamsApi.FindSysParams)       // 根据ID获取参数
		sysParamsRouterWithoutRecord.GET("getSysParamsList", r.sysParamsApi.GetSysParamsList) // 获取参数列表
		sysParamsRouterWithoutRecord.GET("getSysParam", r.sysParamsApi.GetSysParam)           // 根据Key获取参数
	}
}
