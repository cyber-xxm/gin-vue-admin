package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewAuthorityBtnRouter(authorityBtnApi *system.AuthorityBtnApi) *AuthorityBtnRouter {
	return &AuthorityBtnRouter{
		authorityBtnApi: authorityBtnApi,
	}
}

type AuthorityBtnRouter struct {
	authorityBtnApi *system.AuthorityBtnApi
}

func (r *AuthorityBtnRouter) InitAuthorityBtnRouterRouter(router *gin.RouterGroup, recordService *service.OperationRecordService) {
	authorityRouter := router.Group("authorityBtn").Use(middleware.OperationRecord(recordService))
	authorityRouterWithoutRecord := router.Group("authorityBtn")
	{
		authorityRouterWithoutRecord.POST("getAuthorityBtn", r.authorityBtnApi.GetAuthorityBtn)
		authorityRouterWithoutRecord.POST("setAuthorityBtn", r.authorityBtnApi.SetAuthorityBtn)
		authorityRouterWithoutRecord.POST("canRemoveAuthorityBtn", r.authorityBtnApi.CanRemoveAuthorityBtn)
	}
}
