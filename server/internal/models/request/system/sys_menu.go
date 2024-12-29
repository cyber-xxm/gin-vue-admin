package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
)

// AddMenuAuthorityInfo Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []system.SysBaseMenu `json:"menus"`
	AuthorityId uint                 `json:"authorityId"` // 角色ID
}

func DefaultMenu() []system.SysBaseMenu {
	m := system.SysBaseMenu{
		ParentId:  0,
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: system.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}
	m.ID = 1
	return []system.SysBaseMenu{m}
}
