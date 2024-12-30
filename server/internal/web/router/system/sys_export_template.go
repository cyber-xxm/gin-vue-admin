package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewSysExportTemplateRouter(exportTemplateApi *system.SysExportTemplateApi) *SysExportTemplateRouter {
	return &SysExportTemplateRouter{
		exportTemplateApi: exportTemplateApi,
	}
}

type SysExportTemplateRouter struct {
	exportTemplateApi *system.SysExportTemplateApi
}

// InitSysExportTemplateRouter 初始化 导出模板 路由信息
func (r *SysExportTemplateRouter) InitSysExportTemplateRouter(router *gin.RouterGroup, recordService *service.OperationRecordService) {
	sysExportTemplateRouter := router.Group("sysExportTemplate").Use(middleware.OperationRecord(recordService))
	sysExportTemplateRouterWithoutRecord := router.Group("sysExportTemplate")
	{
		sysExportTemplateRouter.POST("createSysExportTemplate", r.exportTemplateApi.CreateSysExportTemplate)             // 新建导出模板
		sysExportTemplateRouter.DELETE("deleteSysExportTemplate", r.exportTemplateApi.DeleteSysExportTemplate)           // 删除导出模板
		sysExportTemplateRouter.DELETE("deleteSysExportTemplateByIds", r.exportTemplateApi.DeleteSysExportTemplateByIds) // 批量删除导出模板
		sysExportTemplateRouter.PUT("updateSysExportTemplate", r.exportTemplateApi.UpdateSysExportTemplate)              // 更新导出模板
		sysExportTemplateRouter.POST("importExcel", r.exportTemplateApi.ImportExcel)                                     // 更新导出模板
	}
	{
		sysExportTemplateRouterWithoutRecord.GET("findSysExportTemplate", r.exportTemplateApi.FindSysExportTemplate)       // 根据ID获取导出模板
		sysExportTemplateRouterWithoutRecord.GET("getSysExportTemplateList", r.exportTemplateApi.GetSysExportTemplateList) // 获取导出模板列表
		sysExportTemplateRouterWithoutRecord.GET("exportExcel", r.exportTemplateApi.ExportExcel)                           // 导出表格
		sysExportTemplateRouterWithoutRecord.GET("exportTemplate", r.exportTemplateApi.ExportTemplate)                     // 导出表格模板
	}
}
