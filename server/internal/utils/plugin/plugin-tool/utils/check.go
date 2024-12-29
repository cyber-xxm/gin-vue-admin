package utils

import (
	"fmt"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"gorm.io/gorm"
)

func RegisterApis(db *gorm.DB, apis ...system.SysApi) {
	var count int64
	var apiPaths []string
	for i := range apis {
		apiPaths = append(apiPaths, apis[i].Path)
	}
	db.Find(&[]system.SysApi{}, "path in (?)", apiPaths).Count(&count)
	if count > 0 {
		return
	}
	err := db.Create(&apis).Error
	if err != nil {
		fmt.Println(err)
	}
}

func RegisterMenus(db *gorm.DB, menus ...system.SysBaseMenu) {
	var count int64
	var menuNames []string
	parentMenu := menus[0]
	otherMenus := menus[1:]
	for i := range menus {
		menuNames = append(menuNames, menus[i].Name)
	}
	db.Find(&[]system.SysBaseMenu{}, "name in (?)", menuNames).Count(&count)
	if count > 0 {
		return
	}
	err := db.Create(&parentMenu).Error
	if err != nil {
		fmt.Println(err)
	}
	for i := range otherMenus {
		pid := parentMenu.ID
		otherMenus[i].ParentId = pid
	}
	err = db.Create(&otherMenus).Error
	if err != nil {
		fmt.Println(err)
	}
}
