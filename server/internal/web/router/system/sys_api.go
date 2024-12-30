package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewSystemApiRouter(systemApi *system.SystemApi) *ApiRouter {
	return &ApiRouter{
		systemApi: systemApi,
	}
}

type ApiRouter struct {
	systemApi *system.SystemApi
}

func (r *ApiRouter) InitApiRouter(router *gin.RouterGroup, pub *gin.RouterGroup, recordService *service.OperationRecordService) {
	apiRouter := router.Group("api").Use(middleware.OperationRecord(recordService))
	apiRouterWithoutRecord := router.Group("api")

	apiPublicRouterWithoutRecord := pub.Group("api")
	{
		apiRouter.GET("getApiGroups", r.systemApi.GetApiGroups)          // 获取路由组
		apiRouter.GET("syncApi", r.systemApi.SyncApi)                    // 同步Api
		apiRouter.POST("ignoreApi", r.systemApi.IgnoreApi)               // 忽略Api
		apiRouter.POST("enterSyncApi", r.systemApi.EnterSyncApi)         // 确认同步Api
		apiRouter.POST("createApi", r.systemApi.CreateApi)               // 创建Api
		apiRouter.POST("deleteApi", r.systemApi.DeleteApi)               // 删除Api
		apiRouter.POST("getApiById", r.systemApi.GetApiById)             // 获取单条Api消息
		apiRouter.POST("updateApi", r.systemApi.UpdateApi)               // 更新api
		apiRouter.DELETE("deleteApisByIds", r.systemApi.DeleteApisByIds) // 删除选中api
	}
	{
		apiRouterWithoutRecord.POST("getAllApis", r.systemApi.GetAllApis) // 获取所有api
		apiRouterWithoutRecord.POST("getApiList", r.systemApi.GetApiList) // 获取Api列表
	}
	{
		apiPublicRouterWithoutRecord.GET("freshCasbin", r.systemApi.FreshCasbin) // 刷新casbin权限
	}
}
