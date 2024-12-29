package task

import (
	"errors"
	"fmt"
	"github.com/cyber-xxm/gin-vue-admin/internal/models"
	"time"

	"gorm.io/gorm"
)

//@author: [songzhibin97](https://github.com/songzhibin97)
//@function: ClearTable
//@description: 清理数据库表数据
//@param: db(数据库对象) *gorm.DB, tableName(表名) string, compareField(比较字段) string, interval(间隔) string
//@return: error

func ClearTable(db *gorm.DB) error {
	var ClearTableDetail []models.ClearDB

	ClearTableDetail = append(ClearTableDetail, models.ClearDB{
		TableName:    "sys_operation_records",
		CompareField: "created_at",
		Interval:     "2160h",
	})

	ClearTableDetail = append(ClearTableDetail, models.ClearDB{
		TableName:    "jwt_blacklists",
		CompareField: "created_at",
		Interval:     "168h",
	})

	if db == nil {
		return errors.New("db Cannot be empty")
	}

	for _, detail := range ClearTableDetail {
		duration, err := time.ParseDuration(detail.Interval)
		if err != nil {
			return err
		}
		if duration < 0 {
			return errors.New("parse duration < 0")
		}
		err = db.Debug().Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", detail.TableName, detail.CompareField), time.Now().Add(-duration)).Error
		if err != nil {
			return err
		}
	}
	return nil
}