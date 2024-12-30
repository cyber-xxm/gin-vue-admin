package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/cyber-xxm/gin-vue-admin/internal/initialize"
	models "github.com/cyber-xxm/gin-vue-admin/internal/models/config"
	zap_logger "github.com/cyber-xxm/gin-vue-admin/internal/utils/zap-logger"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/router"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer(ctx context.Context) error {
	cfg := ctx.Value("config").(*models.Server)
	db := ctx.Value("db").(*gorm.DB)
	if cfg != nil && db != nil {
		return errors.New("系统配置为空，程序退出")
	}
	if cfg.System.UseMultipoint || cfg.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		initialize.RedisList()
	}

	if cfg.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从db加载jwt数据
	system.LoadAll()

	r := router.Routers(ctx)

	address := fmt.Sprintf(":%d", cfg.System.Addr)
	s := initServer(address, r)

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
	** 剔除授权标识需购买商用授权：https://gin-vue-admin.com/empower/index.html **`, address)
	err := s.ListenAndServe()
	if err != nil {
		zap_logger.Error(err.Error())
	}
	return err
}
