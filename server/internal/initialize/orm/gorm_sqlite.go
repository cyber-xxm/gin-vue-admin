package orm

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"path/filepath"
)

type Sqlite struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (s *Sqlite) Dsn() string {
	return filepath.Join(s.Path, s.Dbname+".db")
}

// GormSqlite 初始化Sqlite数据库
func GormSqlite(dbType string, s Sqlite) *gorm.DB {
	if s.Dbname == "" {
		return nil
	}

	if db, err := gorm.Open(sqlite.Open(s.Dsn()), Config(dbType, s.Prefix, s.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}

// GormSqliteByConfig 初始化Sqlite数据库用过传入配置
func GormSqliteByConfig(dbType string, s Sqlite) *gorm.DB {
	if s.Dbname == "" {
		return nil
	}

	if db, err := gorm.Open(sqlite.Open(s.Dsn()), Config(dbType, s.Prefix, s.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}
