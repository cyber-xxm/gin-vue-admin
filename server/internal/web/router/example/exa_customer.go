package example

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct{}

func (r *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("customer")
	{
		customerRouter.POST("customer", r.exaCustomerApi.CreateExaCustomer) // 创建客户
		customerRouter.PUT("customer", r.exaCustomerApi.UpdateExaCustomer)  // 更新客户
		customerRouter.DELETE("customer", exaCustomerApi.DeleteExaCustomer) // 删除客户
	}
	{
		customerRouterWithoutRecord.GET("customer", r.exaCustomerApi.GetExaCustomer)         // 获取单一客户信息
		customerRouterWithoutRecord.GET("customerList", r.exaCustomerApi.GetExaCustomerList) // 获取客户列表
	}
}
