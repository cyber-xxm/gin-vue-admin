package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/gin-gonic/gin"
)

func NewAutoCodeHistoryRouter(autocodeHistoryApi *system.AutoCodeHistoryApi) *AutoCodeHistoryRouter {
	return &AutoCodeHistoryRouter{
		autocodeHistoryApi: autocodeHistoryApi,
	}
}

type AutoCodeHistoryRouter struct {
	autocodeHistoryApi *system.AutoCodeHistoryApi
}

func (r *AutoCodeHistoryRouter) InitAutoCodeHistoryRouter(router *gin.RouterGroup) {
	autoCodeHistoryRouter := router.Group("autoCode")
	{
		autoCodeHistoryRouter.POST("getMeta", r.autocodeHistoryApi.First)         // 根据id获取meta信息
		autoCodeHistoryRouter.POST("rollback", r.autocodeHistoryApi.RollBack)     // 回滚
		autoCodeHistoryRouter.POST("delSysHistory", r.autocodeHistoryApi.Delete)  // 删除回滚记录
		autoCodeHistoryRouter.POST("getSysHistory", r.autocodeHistoryApi.GetList) // 获取回滚记录分页
	}
}
