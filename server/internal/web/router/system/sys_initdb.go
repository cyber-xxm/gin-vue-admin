package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/gin-gonic/gin"
)

func NewInitRouter(dbApi *system.DBApi) *InitRouter {
	return &InitRouter{
		dbApi: dbApi,
	}
}

type InitRouter struct {
	dbApi *system.DBApi
}

func (r *InitRouter) InitInitRouter(router *gin.RouterGroup) {
	initRouter := router.Group("init")
	{
		initRouter.POST("initdb", r.dbApi.InitDB)   // 初始化数据库
		initRouter.POST("checkdb", r.dbApi.CheckDB) // 检测是否需要初始化数据库
	}
}
