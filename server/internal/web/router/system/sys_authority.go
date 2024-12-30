package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewAuthorityRouter(authorityApi *system.AuthorityApi, recordService *service.OperationRecordService) *AuthorityRouter {
	return &AuthorityRouter{
		authorityApi:  authorityApi,
		recordService: recordService,
	}
}

type AuthorityRouter struct {
	authorityApi  *system.AuthorityApi
	recordService *service.OperationRecordService
}

func (r *AuthorityRouter) InitAuthorityRouter(router *gin.RouterGroup) {
	authorityRouter := router.Group("authority").Use(middleware.OperationRecord(r.recordService))
	authorityRouterWithoutRecord := router.Group("authority")
	{
		authorityRouter.POST("createAuthority", r.authorityApi.CreateAuthority)   // 创建角色
		authorityRouter.POST("deleteAuthority", r.authorityApi.DeleteAuthority)   // 删除角色
		authorityRouter.PUT("updateAuthority", r.authorityApi.UpdateAuthority)    // 更新角色
		authorityRouter.POST("copyAuthority", r.authorityApi.CopyAuthority)       // 拷贝角色
		authorityRouter.POST("setDataAuthority", r.authorityApi.SetDataAuthority) // 设置角色资源权限
	}
	{
		authorityRouterWithoutRecord.POST("getAuthorityList", r.authorityApi.GetAuthorityList) // 获取角色列表
	}
}
