package system

import (
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
)

type SysAutoCodePackageCreate struct {
	Desc        string `json:"desc" example:"描述"`
	Label       string `json:"label" example:"展示名"`
	Template    string `json:"template"  example:"模版"`
	PackageName string `json:"packageName" example:"包名"`
	Module      string `json:"-" example:"模块"`
}

func (r *SysAutoCodePackageCreate) AutoCode() AutoCode {
	return AutoCode{
		Package: r.PackageName,
		Module:  global.GVA_CONFIG.AutoCode.Module,
	}
}

func (r *SysAutoCodePackageCreate) Create() system.SysAutoCodePackage {
	return system.SysAutoCodePackage{
		Desc:        r.Desc,
		Label:       r.Label,
		Template:    r.Template,
		PackageName: r.PackageName,
		Module:      global.GVA_CONFIG.AutoCode.Module,
	}
}
