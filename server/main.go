package main

import (
	"context"
	"github.com/cyber-xxm/gin-vue-admin/internal/initialize"
	"github.com/cyber-xxm/gin-vue-admin/internal/initialize/config"
	"github.com/cyber-xxm/gin-vue-admin/internal/initialize/orm"
	models "github.com/cyber-xxm/gin-vue-admin/internal/models/config"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/example"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils/zap-logger"
	"github.com/cyber-xxm/gin-vue-admin/internal/web"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// 这部分 @Tag 设置用于排序, 需要排序的接口请按照下面的格式添加
// swag init 对 @Tag 只会从入口文件解析, 默认 main.go
// 也可通过 --generalInfo flag 指定其他文件
// @Tag.Name        Base
// @Tag.Name        SysUser
// @Tag.Description 用户

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.7.8-beta1
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	ctx := context.WithoutCancel(context.Background())
	viper, cfg := config.Viper()
	ctx = context.WithValue(ctx, "config", cfg)
	ctx = context.WithValue(ctx, "viper", viper)
	initialize.OtherInit(ctx)
	logger := zap_logger.Zap(cfg.Zap.Director, cfg.Zap.Levels(), cfg.Zap.ShowLine)
	zap_logger.NewZapLogger(logger) // 初始化zap日志库
	zap.ReplaceGlobals(logger)
	ormDb := Gorm(cfg.System.DbType, cfg) // gorm连接数据库
	ctx = context.WithValue(ctx, "db", ormDb)
	initialize.Timer(ormDb)
	initialize.DBList(cfg.DBList)
	if ormDb != nil {
		RegisterTables(ormDb) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := ormDb.DB()
		defer db.Close()
	}
	err := web.RunWindowsServer(ctx)
	if err != nil {
		logger.Error(err.Error())
	}
}

func Gorm(dbType string, cfg models.Server) *gorm.DB {
	switch dbType {
	case "mysql":
		return orm.GormMysql(dbType, orm.Mysql{GeneralDB: cfg.Mysql.GeneralDB})
	case "pgsql":
		return orm.GormPgSql(dbType, orm.Pgsql{GeneralDB: cfg.Pgsql.GeneralDB})
	case "oracle":
		return orm.GormOracle(dbType, orm.Oracle{GeneralDB: cfg.Oracle.GeneralDB})
	case "mssql":
		return orm.GormMssql(dbType, orm.Mssql{GeneralDB: cfg.Mssql.GeneralDB})
	case "sqlite":
		return orm.GormSqlite(dbType, orm.Sqlite{GeneralDB: cfg.Sqlite.GeneralDB})
	default:
		return orm.GormMysql("mysql", orm.Mysql{GeneralDB: cfg.Mysql.GeneralDB})
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.SysApi{},
		system.SysIgnoreApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCodePackage{},
		system.SysExportTemplate{},
		system.Condition{},
		system.JoinTemplate{},
		system.SysParams{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	)
	if err != nil {
		zap_logger.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	err = bizModel(db)

	if err != nil {
		zap_logger.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	zap_logger.Info("register table success")
}

func bizModel(db *gorm.DB) error {
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
