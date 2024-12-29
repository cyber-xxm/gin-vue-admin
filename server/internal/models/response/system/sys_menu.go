package system

import (
	system2 "github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
)

type SysMenusResponse struct {
	Menus []system2.SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []system2.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu system2.SysBaseMenu `json:"menu"`
}
