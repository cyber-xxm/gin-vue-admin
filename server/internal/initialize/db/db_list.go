package db

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/initialize/db/orm"
	"gorm.io/gorm"
)

const sys = "system"

type DsnProvider interface {
	Dsn() string
}

func DBList(dbList []orm.SpecializedDB) map[string]*gorm.DB {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range dbList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = orm.GormMysqlByConfig(info.Type, orm.Mysql{GeneralDB: info.GeneralDB})
		case "mssql":
			dbMap[info.AliasName] = orm.GormMssqlByConfig(info.Type, orm.Mssql{GeneralDB: info.GeneralDB})
		case "pgsql":
			dbMap[info.AliasName] = orm.GormPgSqlByConfig(info.Type, orm.Pgsql{GeneralDB: info.GeneralDB})
		case "oracle":
			dbMap[info.AliasName] = orm.GormOracleByConfig(info.Type, orm.Oracle{GeneralDB: info.GeneralDB})
		default:
			continue
		}
	}
	// 做特殊判断,是否有迁移
	// 适配低版本迁移多数据库版本
	//if sysDB, ok := dbMap[sys]; ok {
	//	global.GVA_DB = sysDB
	//}
	return dbMap
}
