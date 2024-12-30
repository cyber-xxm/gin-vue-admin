package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewConfigRouter(configApi *system.ConfigApi) *ConfigRouter {
	return &ConfigRouter{
		configApi: configApi,
	}
}

type ConfigRouter struct {
	configApi *system.ConfigApi
}

func (r *ConfigRouter) InitSystemRouter(router *gin.RouterGroup, recordService *service.OperationRecordService) {
	sysRouter := router.Group("system").Use(middleware.OperationRecord(recordService))
	sysRouterWithoutRecord := router.Group("system")

	{
		sysRouter.POST("setSystemConfig", r.configApi.SetSystemConfig) // 设置配置文件内容
		sysRouter.POST("reloadSystem", r.configApi.ReloadSystem)       // 重启服务
	}
	{
		sysRouterWithoutRecord.POST("getSystemConfig", r.configApi.GetSystemConfig) // 获取配置文件内容
		sysRouterWithoutRecord.POST("getServerInfo", r.configApi.GetServerInfo)     // 获取服务器信息
	}
}
