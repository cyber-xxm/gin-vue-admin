package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewDictionaryDetailRouter(dictionaryDetailApi *system.DictionaryDetailApi, recordService *service.OperationRecordService) *DictionaryDetailRouter {
	return &DictionaryDetailRouter{
		dictionaryDetailApi: dictionaryDetailApi,
	}
}

type DictionaryDetailRouter struct {
	dictionaryDetailApi *system.DictionaryDetailApi
	recordService       *service.OperationRecordService
}

func (r *DictionaryDetailRouter) InitSysDictionaryDetailRouter(router *gin.RouterGroup) {
	dictionaryDetailRouter := router.Group("sysDictionaryDetail").Use(middleware.OperationRecord(r.recordService))
	dictionaryDetailRouterWithoutRecord := router.Group("sysDictionaryDetail")
	{
		dictionaryDetailRouter.POST("createSysDictionaryDetail", r.dictionaryDetailApi.CreateSysDictionaryDetail)   // 新建SysDictionaryDetail
		dictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", r.dictionaryDetailApi.DeleteSysDictionaryDetail) // 删除SysDictionaryDetail
		dictionaryDetailRouter.PUT("updateSysDictionaryDetail", r.dictionaryDetailApi.UpdateSysDictionaryDetail)    // 更新SysDictionaryDetail
	}
	{
		dictionaryDetailRouterWithoutRecord.GET("findSysDictionaryDetail", r.dictionaryDetailApi.FindSysDictionaryDetail)       // 根据ID获取SysDictionaryDetail
		dictionaryDetailRouterWithoutRecord.GET("getSysDictionaryDetailList", r.dictionaryDetailApi.GetSysDictionaryDetailList) // 获取SysDictionaryDetail列表
	}
}
