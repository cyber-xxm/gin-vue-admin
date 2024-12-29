package web

import (
	"fmt"
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/initialize/db"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/router"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		db.Redis()
		db.RedisList()
	}

	if global.GVA_CONFIG.System.UseMongo {
		err := db.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := router.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	zap_logger.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 gin-vue-admin
	当前版本:v2.7.8-beta1
	加群方式:微信号：shouzi_1994 QQ群：470239250
	项目地址：https://github.com/cyber-xxm/gin-vue-admin
	插件市场:https://plugin.gin-vue-admin.com
	GVA讨论社区:https://support.qq.com/products/371961
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	--------------------------------------版权声明--------------------------------------
	** 版权所有方：flipped-aurora开源团队 **
	** 版权持有公司：北京翻转极光科技有限责任公司 **
	** 剔除授权标识需购买商用授权：https://gin-vue-admin.com/empower/index.html **
`, address)
	zap_logger.Error(s.ListenAndServe().Error())
}
