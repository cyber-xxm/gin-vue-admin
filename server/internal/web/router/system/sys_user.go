package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(baseApi *system.BaseApi, recordService *service.OperationRecordService) *UserRouter {
	return &UserRouter{
		baseApi:       baseApi,
		recordService: recordService,
	}
}

type UserRouter struct {
	baseApi       *system.BaseApi
	recordService *service.OperationRecordService
}

func (r *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user").Use(middleware.OperationRecord(r.recordService))
	userRouterWithoutRecord := router.Group("user")
	{
		userRouter.POST("admin_register", r.baseApi.Register)               // 管理员注册账号
		userRouter.POST("changePassword", r.baseApi.ChangePassword)         // 用户修改密码
		userRouter.POST("setUserAuthority", r.baseApi.SetUserAuthority)     // 设置用户权限
		userRouter.DELETE("deleteUser", r.baseApi.DeleteUser)               // 删除用户
		userRouter.PUT("setUserInfo", r.baseApi.SetUserInfo)                // 设置用户信息
		userRouter.PUT("setSelfInfo", r.baseApi.SetSelfInfo)                // 设置自身信息
		userRouter.POST("setUserAuthorities", r.baseApi.SetUserAuthorities) // 设置用户权限组
		userRouter.POST("resetPassword", r.baseApi.ResetPassword)           // 设置用户权限组
		userRouter.PUT("setSelfSetting", r.baseApi.SetSelfSetting)          // 用户界面配置
	}
	{
		userRouterWithoutRecord.POST("getUserList", r.baseApi.GetUserList) // 分页获取用户列表
		userRouterWithoutRecord.GET("getUserInfo", r.baseApi.GetUserInfo)  // 获取自身信息
	}
}
