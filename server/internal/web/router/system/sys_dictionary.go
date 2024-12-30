package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewDictionaryRouter(dictionaryApi *system.DictionaryApi, recordService *service.OperationRecordService) *DictionaryRouter {
	return &DictionaryRouter{
		dictionaryApi: dictionaryApi,
	}
}

type DictionaryRouter struct {
	dictionaryApi *system.DictionaryApi
	recordService *service.OperationRecordService
}

func (r *DictionaryRouter) InitSysDictionaryRouter(router *gin.RouterGroup) {
	sysDictionaryRouter := router.Group("sysDictionary").Use(middleware.OperationRecord(r.recordService))
	sysDictionaryRouterWithoutRecord := router.Group("sysDictionary")
	{
		sysDictionaryRouter.POST("createSysDictionary", r.dictionaryApi.CreateSysDictionary)   // 新建SysDictionary
		sysDictionaryRouter.DELETE("deleteSysDictionary", r.dictionaryApi.DeleteSysDictionary) // 删除SysDictionary
		sysDictionaryRouter.PUT("updateSysDictionary", r.dictionaryApi.UpdateSysDictionary)    // 更新SysDictionary
	}
	{
		sysDictionaryRouterWithoutRecord.GET("findSysDictionary", r.dictionaryApi.FindSysDictionary)       // 根据ID获取SysDictionary
		sysDictionaryRouterWithoutRecord.GET("getSysDictionaryList", r.dictionaryApi.GetSysDictionaryList) // 获取SysDictionary列表
	}
}
