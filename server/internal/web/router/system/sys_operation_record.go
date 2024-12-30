package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/gin-gonic/gin"
)

func NewOperationRecordRouter(operationRecordApi *system.OperationRecordApi) *OperationRecordRouter {
	return &OperationRecordRouter{
		operationRecordApi: operationRecordApi,
	}
}

type OperationRecordRouter struct {
	operationRecordApi *system.OperationRecordApi
}

func (r *OperationRecordRouter) InitSysOperationRecordRouter(router *gin.RouterGroup) {
	operationRecordRouter := router.Group("sysOperationRecord")
	{
		operationRecordRouter.POST("createSysOperationRecord", r.operationRecordApi.CreateSysOperationRecord)             // 新建SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecord", r.operationRecordApi.DeleteSysOperationRecord)           // 删除SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecordByIds", r.operationRecordApi.DeleteSysOperationRecordByIds) // 批量删除SysOperationRecord
		operationRecordRouter.GET("findSysOperationRecord", r.operationRecordApi.FindSysOperationRecord)                  // 根据ID获取SysOperationRecord
		operationRecordRouter.GET("getSysOperationRecordList", r.operationRecordApi.GetSysOperationRecordList)            // 获取SysOperationRecord列表

	}
}
