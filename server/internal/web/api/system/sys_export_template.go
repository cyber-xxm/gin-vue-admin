package system

import (
	"fmt"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request"
	system2 "github.com/cyber-xxm/gin-vue-admin/internal/models/request/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/response"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"
	zap_logger "github.com/cyber-xxm/gin-vue-admin/internal/utils/zap-logger"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewSysExportTemplateApi(db *gorm.DB) *SysExportTemplateApi {
	return &SysExportTemplateApi{
		SysExportTemplateService: service.NewSysExportTemplateService(db),
	}
}

type SysExportTemplateApi struct {
	SysExportTemplateService *service.SysExportTemplateService
}

// CreateSysExportTemplate 创建导出模板
// @Tags SysExportTemplate
// @Summary 创建导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysExportTemplate true "创建导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysExportTemplate/createSysExportTemplate [post]
func (a *SysExportTemplateApi) CreateSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindJSON(&sysExportTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(sysExportTemplate, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := a.SysExportTemplateService.CreateSysExportTemplate(&sysExportTemplate); err != nil {
		zap_logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysExportTemplate 删除导出模板
// @Tags SysExportTemplate
// @Summary 删除导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysExportTemplate true "删除导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysExportTemplate/deleteSysExportTemplate [delete]
func (a *SysExportTemplateApi) DeleteSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindJSON(&sysExportTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := a.SysExportTemplateService.DeleteSysExportTemplate(sysExportTemplate); err != nil {
		zap_logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysExportTemplateByIds 批量删除导出模板
// @Tags SysExportTemplate
// @Summary 批量删除导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysExportTemplate/deleteSysExportTemplateByIds [delete]
func (a *SysExportTemplateApi) DeleteSysExportTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := a.SysExportTemplateService.DeleteSysExportTemplateByIds(IDS); err != nil {
		zap_logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysExportTemplate 更新导出模板
// @Tags SysExportTemplate
// @Summary 更新导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysExportTemplate true "更新导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysExportTemplate/updateSysExportTemplate [put]
func (a *SysExportTemplateApi) UpdateSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindJSON(&sysExportTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(sysExportTemplate, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := a.SysExportTemplateService.UpdateSysExportTemplate(sysExportTemplate); err != nil {
		zap_logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysExportTemplate 用id查询导出模板
// @Tags SysExportTemplate
// @Summary 用id查询导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysExportTemplate true "用id查询导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysExportTemplate/findSysExportTemplate [get]
func (a *SysExportTemplateApi) FindSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindQuery(&sysExportTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resysExportTemplate, err := a.SysExportTemplateService.GetSysExportTemplate(sysExportTemplate.ID); err != nil {
		zap_logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysExportTemplate": resysExportTemplate}, c)
	}
}

// GetSysExportTemplateList 分页获取导出模板列表
// @Tags SysExportTemplate
// @Summary 分页获取导出模板列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query systemReq.SysExportTemplateSearch true "分页获取导出模板列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysExportTemplate/getSysExportTemplateList [get]
func (a *SysExportTemplateApi) GetSysExportTemplateList(c *gin.Context) {
	var pageInfo system2.SysExportTemplateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := a.SysExportTemplateService.GetSysExportTemplateInfoList(pageInfo); err != nil {
		zap_logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// ExportExcel 导出表格
// @Tags SysExportTemplate
// @Summary 导出表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/exportExcel [get]
func (a *SysExportTemplateApi) ExportExcel(c *gin.Context) {
	templateID := c.Query("templateID")
	queryParams := c.Request.URL.Query()
	if templateID == "" {
		response.FailWithMessage("模板ID不能为空", c)
		return
	}
	if file, name, err := a.SysExportTemplateService.ExportExcel(templateID, queryParams); err != nil {
		zap_logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name+utils.RandomString(6)+".xlsx")) // 对下载的文件重命名
		c.Header("success", "true")
		c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", file.Bytes())
	}
}

// ExportTemplate 导出表格模板
// @Tags SysExportTemplate
// @Summary 导出表格模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/ExportTemplate [get]
func (a *SysExportTemplateApi) ExportTemplate(c *gin.Context) {
	templateID := c.Query("templateID")
	if templateID == "" {
		response.FailWithMessage("模板ID不能为空", c)
		return
	}
	if file, name, err := a.SysExportTemplateService.ExportTemplate(templateID); err != nil {
		zap_logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name+"模板.xlsx")) // 对下载的文件重命名
		c.Header("success", "true")
		c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", file.Bytes())
	}
}

// ImportExcel 导入表格
// @Tags SysImportTemplate
// @Summary 导入表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/importExcel [post]
func (a *SysExportTemplateApi) ImportExcel(c *gin.Context) {
	templateID := c.Query("templateID")
	if templateID == "" {
		response.FailWithMessage("模板ID不能为空", c)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		zap_logger.Error("文件获取失败!", zap.Error(err))
		response.FailWithMessage("文件获取失败", c)
		return
	}
	if err := a.SysExportTemplateService.ImportExcel(templateID, file); err != nil {
		zap_logger.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("导入成功", c)
	}
}
