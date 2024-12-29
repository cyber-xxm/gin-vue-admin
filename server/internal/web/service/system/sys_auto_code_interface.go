package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/initialize/db/orm"
	sysResp "github.com/cyber-xxm/gin-vue-admin/internal/models/response/system"
	"gorm.io/gorm"
)

func NewAutoCodeInterfaceService(db *gorm.DB) *AutoCodeInterfaceService {
	return &AutoCodeInterfaceService{
		db: db,
	}
}

type AutoCodeInterfaceService struct {
	db *gorm.DB
}

type Database interface {
	GetDB(businessDB string) (data []sysResp.Db, err error)
	GetTables(businessDB string, dbName string) (data []sysResp.Table, err error)
	GetColumn(businessDB string, tableName string, dbName string) (data []sysResp.Column, err error)
}

func (s *AutoCodeInterfaceService) Database(dbType, businessDB string, dbList []orm.SpecializedDB) Database {

	if businessDB == "" {
		switch dbType {
		case "mysql":
			return NewAutoCodeMysqlService(s.db)
		case "pgsql":
			return NewAutoCodePgsqlService(s.db)
		case "mssql":
			return NewAutoCodeMssqlService(s.db)
		case "oracle":
			return NewAutoCodeOracleService(s.db)
		case "sqlite":
			return NewAutoCodeSqliteService(s.db)
		default:
			return NewAutoCodeMssqlService(s.db)
		}
	} else {
		for _, info := range dbList {
			if info.AliasName == businessDB {
				switch info.Type {
				case "mysql":
					return NewAutoCodeMysqlService(s.db)
				case "mssql":
					return NewAutoCodeMssqlService(s.db)
				case "pgsql":
					return NewAutoCodePgsqlService(s.db)
				case "oracle":
					return NewAutoCodeOracleService(s.db)
				case "sqlite":
					return NewAutoCodeSqliteService(s.db)
				default:
					return NewAutoCodeMssqlService(s.db)
				}
			}
		}
		return NewAutoCodeMysqlService(s.db)
	}

}
