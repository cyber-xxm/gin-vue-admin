package system

import (
	"fmt"
	"github.com/cyber-xxm/gin-vue-admin/global"
	sysResp "github.com/cyber-xxm/gin-vue-admin/internal/models/response/system"
	"gorm.io/gorm"
	"path/filepath"
	"strings"
)

func NewAutoCodeSqliteService(db *gorm.DB) *AutoCodeSqliteService {
	return &AutoCodeSqliteService{
		db: db,
	}
}

type AutoCodeSqliteService struct {
	db *gorm.DB
}

// GetDB 获取数据库的所有数据库名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *AutoCodeSqliteService) GetDB(businessDB string) (data []sysResp.Db, err error) {
	var entities []sysResp.Db
	sql := "PRAGMA database_list;"
	var databaseList []struct {
		File string `gorm:"column:file"`
	}
	if businessDB == "" {
		err = s.db.Raw(sql).Find(&databaseList).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Find(&databaseList).Error
	}
	for _, database := range databaseList {
		if database.File != "" {
			fileName := filepath.Base(database.File)
			fileExt := filepath.Ext(fileName)
			fileNameWithoutExt := strings.TrimSuffix(fileName, fileExt)

			entities = append(entities, sysResp.Db{fileNameWithoutExt})
		}
	}
	// entities = append(entities, system.Db{global.GVA_CONFIG.Sqlite.Dbname})
	return entities, err
}

// GetTables 获取数据库的所有表名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *AutoCodeSqliteService) GetTables(businessDB string, dbName string) (data []sysResp.Table, err error) {
	var entities []sysResp.Table
	sql := `SELECT name FROM sqlite_master WHERE type='table'`
	tabelNames := []string{}
	if businessDB == "" {
		err = s.db.Raw(sql).Find(&tabelNames).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Find(&tabelNames).Error
	}
	for _, tabelName := range tabelNames {
		entities = append(entities, sysResp.Table{tabelName})
	}
	return entities, err
}

// GetColumn 获取指定数据表的所有字段名,类型值等
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *AutoCodeSqliteService) GetColumn(businessDB string, tableName string, dbName string) (data []sysResp.Column, err error) {
	var entities []sysResp.Column
	sql := fmt.Sprintf("PRAGMA table_info(%s);", tableName)
	var columnInfos []struct {
		Name string `gorm:"column:name"`
		Type string `gorm:"column:type"`
		Pk   int    `gorm:"column:pk"`
	}
	if businessDB == "" {
		err = s.db.Raw(sql).Scan(&columnInfos).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Scan(&columnInfos).Error
	}
	for _, columnInfo := range columnInfos {
		entities = append(entities, sysResp.Column{
			ColumnName: columnInfo.Name,
			DataType:   columnInfo.Type,
			PrimaryKey: columnInfo.Pk == 1,
		})
	}
	return entities, err
}
