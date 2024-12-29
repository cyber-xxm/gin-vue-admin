package initialize

import (
	"context"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils/plugin/plugin-tool/utils"
	"gorm.io/gorm"
)

func Menu(ctx context.Context) {
	db := ctx.Value("db").(*gorm.DB)
	entities := []system.SysBaseMenu{
		{
			ParentId:  24,
			Path:      "anInfo",
			Name:      "anInfo",
			Hidden:    false,
			Component: "plugin/announcement/view/info.vue",
			Sort:      5,
			Meta:      system.Meta{Title: "公告管理", Icon: "box"},
		},
	}
	utils.RegisterMenus(db, entities...)
}
