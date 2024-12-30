package orm

/*
 * @Author: 逆光飞翔 191180776@qq.com
 * @Date: 2022-12-08 17:25:49
 * @LastEditors: 逆光飞翔 191180776@qq.com
 * @LastEditTime: 2022-12-08 18:00:00
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Mssql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// Dsn "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
func (m *Mssql) Dsn() string {
	return "sqlserver://" + m.Username + ":" + m.Password + "@" + m.Path + ":" + m.Port + "?database=" + m.Dbname + "&encrypt=disable"
}

// GormMssql 初始化Mssql数据库
// Author [LouisZhang](191180776@qq.com)
func GormMssql(dbType string, m Mssql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), Config(dbType, m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormMssqlByConfig 初始化Mysql数据库用过传入配置
func GormMssqlByConfig(dbType string, m Mssql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string 类型字段的默认长度
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), Config(dbType, m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
