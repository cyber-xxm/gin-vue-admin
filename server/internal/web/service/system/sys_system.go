package system

import (
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/config"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewSystemConfigService(db *gorm.DB) *SystemConfigService {
	return &SystemConfigService{
		db: db,
	}
}

type SystemConfigService struct {
	db *gorm.DB
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: conf config.Server, err error

func (s *SystemConfigService) GetSystemConfig() (conf models.Server, err error) {
	return global.GVA_CONFIG, nil
}

// @description   set system config,
//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetSystemConfig
//@description: 设置配置文件
//@param: system model.System
//@return: err error

func (s *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetServerInfo
//@description: 获取服务器信息
//@return: server *utils.Server, err error

func (s *SystemConfigService) GetServerInfo() (*utils.Server, error) {
	var server utils.Server
	var err error
	server.Os = utils.InitOS()
	if server.Cpu, err = utils.InitCPU(); err != nil {
		zap_logger.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &server, err
	}
	if server.Ram, err = utils.InitRAM(); err != nil {
		zap_logger.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &server, err
	}
	if server.Disk, err = utils.InitDisk(); err != nil {
		zap_logger.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &server, err
	}

	return &server, nil
}
