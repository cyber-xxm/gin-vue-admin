package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
)

func NewMenuRouter(authorityMenuApi *system.MenuApi, recordService *service.OperationRecordService) *MenuRouter {
	return &MenuRouter{
		authorityMenuApi: authorityMenuApi,
	}
}

type MenuRouter struct {
	authorityMenuApi *system.MenuApi
	recordService    *service.OperationRecordService
}

func (r *MenuRouter) InitMenuRouter(router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := router.Group("menu").Use(middleware.OperationRecord(r.recordService))
	menuRouterWithoutRecord := router.Group("menu")
	{
		menuRouter.POST("addBaseMenu", r.authorityMenuApi.AddBaseMenu)           // 新增菜单
		menuRouter.POST("addMenuAuthority", r.authorityMenuApi.AddMenuAuthority) //	增加menu和角色关联关系
		menuRouter.POST("deleteBaseMenu", r.authorityMenuApi.DeleteBaseMenu)     // 删除菜单
		menuRouter.POST("updateBaseMenu", r.authorityMenuApi.UpdateBaseMenu)     // 更新菜单
	}
	{
		menuRouterWithoutRecord.POST("getMenu", r.authorityMenuApi.GetMenu)                   // 获取菜单树
		menuRouterWithoutRecord.POST("getMenuList", r.authorityMenuApi.GetMenuList)           // 分页获取基础menu列表
		menuRouterWithoutRecord.POST("getBaseMenuTree", r.authorityMenuApi.GetBaseMenuTree)   // 获取用户动态路由
		menuRouterWithoutRecord.POST("getMenuAuthority", r.authorityMenuApi.GetMenuAuthority) // 获取指定角色menu
		menuRouterWithoutRecord.POST("getBaseMenuById", r.authorityMenuApi.GetBaseMenuById)   // 根据id获取菜单
	}
	return menuRouter
}
