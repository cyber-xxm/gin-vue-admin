package system

import "github.com/cyber-xxm/gin-vue-admin/internal/models/request"

type SysBaseMenuBtn struct {
	request.CommonModel
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:菜单ID"`
}
