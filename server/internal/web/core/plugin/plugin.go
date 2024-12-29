package plugin

import (
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/gin-gonic/gin"
)

func InstallPlugin(PrivateGroup *gin.RouterGroup, PublicRouter *gin.RouterGroup, engine *gin.Engine) {
	if global.GVA_DB == nil {
		zap_logger.Info("项目暂未初始化，无法安装插件，初始化后重启项目即可完成插件安装")
		return
	}
	bizPluginV1(PrivateGroup, PublicRouter)
	bizPluginV2(engine)
}
