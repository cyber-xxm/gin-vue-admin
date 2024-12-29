package initialize

import (
	"context"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils/plugin/plugin-tool/utils"
	"gorm.io/gorm"
)

func Api(ctx context.Context) {
	db := ctx.Value("db").(*gorm.DB)
	entities := []system.SysApi{
		{
			Path:        "/info/createInfo",
			Description: "新建公告",
			ApiGroup:    "公告",
			Method:      "POST",
		},
		{
			Path:        "/info/deleteInfo",
			Description: "删除公告",
			ApiGroup:    "公告",
			Method:      "DELETE",
		},
		{
			Path:        "/info/deleteInfoByIds",
			Description: "批量删除公告",
			ApiGroup:    "公告",
			Method:      "DELETE",
		},
		{
			Path:        "/info/updateInfo",
			Description: "更新公告",
			ApiGroup:    "公告",
			Method:      "PUT",
		},
		{
			Path:        "/info/findInfo",
			Description: "根据ID获取公告",
			ApiGroup:    "公告",
			Method:      "GET",
		},
		{
			Path:        "/info/getInfoList",
			Description: "获取公告列表",
			ApiGroup:    "公告",
			Method:      "GET",
		},
	}
	utils.RegisterApis(db, entities...)
}
