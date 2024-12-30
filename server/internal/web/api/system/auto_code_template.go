package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/response"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"
	zap_logger "github.com/cyber-xxm/gin-vue-admin/internal/utils/zap-logger"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewAutoCodeTemplateApi(db *gorm.DB) *AutoCodeTemplateApi {
	return &AutoCodeTemplateApi{
		AutoCodeTemplateService: service.NewAutoCodeTemplateService(db),
	}
}

type AutoCodeTemplateApi struct {
	AutoCodeTemplateService *service.AutoCodeTemplateService
}

// Preview
// @Tags      AutoCodeTemplate
// @Summary   预览创建后的代码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoCode                                      true  "预览创建代码"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "预览创建后的代码"
// @Router    /autoCode/preview [post]
func (a *AutoCodeTemplateApi) Preview(c *gin.Context) {
	var info system.AutoCode
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(info, utils.AutoCodeVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = info.Pretreatment()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info.PackageT = utils.FirstUpper(info.Package)
	autoCode, err := a.AutoCodeTemplateService.Preview(c.Request.Context(), info)
	if err != nil {
		zap_logger.Error(err.Error(), zap.Error(err))
		response.FailWithMessage("预览失败:"+err.Error(), c)
	} else {
		response.OkWithDetailed(gin.H{"autoCode": autoCode}, "预览成功", c)
	}
}

// Create
// @Tags      AutoCodeTemplate
// @Summary   自动代码模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoCode  true  "创建自动代码"
// @Success   200   {string}  string                 "{"success":true,"data":{},"msg":"创建成功"}"
// @Router    /autoCode/createTemp [post]
func (a *AutoCodeTemplateApi) Create(c *gin.Context) {
	var info system.AutoCode
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(info, utils.AutoCodeVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = info.Pretreatment()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.AutoCodeTemplateService.Create(c.Request.Context(), info)
	if err != nil {
		zap_logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// AddFunc
// @Tags      AddFunc
// @Summary   增加方法
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoCode  true  "增加方法"
// @Success   200   {string}  string                 "{"success":true,"data":{},"msg":"创建成功"}"
// @Router    /autoCode/addFunc [post]
func (a *AutoCodeTemplateApi) AddFunc(c *gin.Context) {
	var info system.AutoFunc
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var tempMap map[string]string
	if info.IsPreview {
		info.Router = "填充router"
		info.FuncName = "填充funcName"
		info.Method = "填充method"
		info.Description = "填充description"
		tempMap, err = a.AutoCodeTemplateService.GetApiAndServer(info)
	} else {
		err = a.AutoCodeTemplateService.AddFunc(info)
	}
	if err != nil {
		zap_logger.Error("注入失败!", zap.Error(err))
		response.FailWithMessage("注入失败", c)
	} else {
		if info.IsPreview {
			response.OkWithDetailed(tempMap, "注入成功", c)
			return
		}
		response.OkWithMessage("注入成功", c)
	}
}
