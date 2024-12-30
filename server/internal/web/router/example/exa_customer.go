package example

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/example"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewCustomerRouter(exaCustomerApi *example.CustomerApi, recordService *service.OperationRecordService) *CustomerRouter {
	return &CustomerRouter{
		exaCustomerApi: exaCustomerApi,
		recordService:  recordService,
	}
}

type CustomerRouter struct {
	exaCustomerApi *example.CustomerApi
	recordService  *service.OperationRecordService
}

func (r *CustomerRouter) InitCustomerRouter(router *gin.RouterGroup) {
	customerRouter := router.Group("customer").Use(middleware.OperationRecord(r.recordService))
	customerRouterWithoutRecord := router.Group("customer")
	{
		customerRouter.POST("customer", r.exaCustomerApi.CreateExaCustomer)   // 创建客户
		customerRouter.PUT("customer", r.exaCustomerApi.UpdateExaCustomer)    // 更新客户
		customerRouter.DELETE("customer", r.exaCustomerApi.DeleteExaCustomer) // 删除客户
	}
	{
		customerRouterWithoutRecord.GET("customer", r.exaCustomerApi.GetExaCustomer)         // 获取单一客户信息
		customerRouterWithoutRecord.GET("customerList", r.exaCustomerApi.GetExaCustomerList) // 获取客户列表
	}
}
