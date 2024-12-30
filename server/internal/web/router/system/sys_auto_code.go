package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/gin-gonic/gin"
)

func NewAutoCodeRouter(autoCodeApi *system.AutoCodeApi,
	autoCodeTemplateApi *system.AutoCodeTemplateApi,
	autoCodePackageApi *system.AutoCodePackageApi,
	autoCodePluginApi *system.AutoCodePluginApi) *AutoCodeRouter {
	return &AutoCodeRouter{
		autoCodeApi:         autoCodeApi,
		autoCodeTemplateApi: autoCodeTemplateApi,
		autoCodePackageApi:  autoCodePackageApi,
		autoCodePluginApi:   autoCodePluginApi,
	}
}

type AutoCodeRouter struct {
	autoCodeApi         *system.AutoCodeApi
	autoCodeTemplateApi *system.AutoCodeTemplateApi
	autoCodePackageApi  *system.AutoCodePackageApi
	autoCodePluginApi   *system.AutoCodePluginApi
}

func (r *AutoCodeRouter) InitAutoCodeRouter(router *gin.RouterGroup, pub *gin.RouterGroup) {
	autoCodeRouter := router.Group("autoCode")
	publicAutoCodeRouter := pub.Group("autoCode")
	{
		autoCodeRouter.GET("getDB", r.autoCodeApi.GetDB)         // 获取数据库
		autoCodeRouter.GET("getTables", r.autoCodeApi.GetTables) // 获取对应数据库的表
		autoCodeRouter.GET("getColumn", r.autoCodeApi.GetColumn) // 获取指定表所有字段信息
	}
	{
		autoCodeRouter.POST("preview", r.autoCodeTemplateApi.Preview)   // 获取自动创建代码预览
		autoCodeRouter.POST("createTemp", r.autoCodeTemplateApi.Create) // 创建自动化代码
		autoCodeRouter.POST("addFunc", r.autoCodeTemplateApi.AddFunc)   // 为代码插入方法
	}
	{
		autoCodeRouter.POST("getPackage", r.autoCodePackageApi.All)       // 获取package包
		autoCodeRouter.POST("delPackage", r.autoCodePackageApi.Delete)    // 删除package包
		autoCodeRouter.POST("createPackage", r.autoCodePackageApi.Create) // 创建package包
	}
	{
		autoCodeRouter.GET("getTemplates", r.autoCodePackageApi.Templates) // 创建package包
	}
	{
		autoCodeRouter.POST("pubPlug", r.autoCodePluginApi.Packaged)      // 打包插件
		autoCodeRouter.POST("installPlugin", r.autoCodePluginApi.Install) // 自动安装插件

	}
	{
		publicAutoCodeRouter.POST("llmAuto", r.autoCodeApi.LLMAuto)
		publicAutoCodeRouter.POST("initMenu", r.autoCodePluginApi.InitMenu) // 同步插件菜单
		publicAutoCodeRouter.POST("initAPI", r.autoCodePluginApi.InitAPI)   // 同步插件API
	}
}
