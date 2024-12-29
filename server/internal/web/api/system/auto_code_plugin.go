package system

import (
	"fmt"
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/response"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewAutoCodePluginApi(db *gorm.DB) *AutoCodePluginApi {
	return &AutoCodePluginApi{
		autoCodePluginService: service.NewAutoCodePluginService(db),
	}
}

type AutoCodePluginApi struct {
	autoCodePluginService service.AutoCodePluginService
}

// Install
// @Tags      AutoCodePlugin
// @Summary   安装插件
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     plug  formData  file                                              true  "this is a test file"
// @Success   200   {object}  response.Response{data=[]interface{},msg=string}  "安装插件成功"
// @Router    /autoCode/installPlugin [post]
func (a *AutoCodePluginApi) Install(c *gin.Context) {
	header, err := c.FormFile("plug")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	web, server, err := a.autoCodePluginService.Install(header)
	webStr := "web插件安装成功"
	serverStr := "server插件安装成功"
	if web == -1 {
		webStr = "web端插件未成功安装，请按照文档自行解压安装，如果为纯后端插件请忽略此条提示"
	}
	if server == -1 {
		serverStr = "server端插件未成功安装，请按照文档自行解压安装，如果为纯前端插件请忽略此条提示"
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData([]interface{}{
		gin.H{
			"code": web,
			"msg":  webStr,
		},
		gin.H{
			"code": server,
			"msg":  serverStr,
		}}, c)
}

// Packaged
// @Tags      AutoCodePlugin
// @Summary   打包插件
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     plugName  query    string  true  "插件名称"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "打包插件成功"
// @Router    /autoCode/pubPlug [post]
func (a *AutoCodePluginApi) Packaged(c *gin.Context) {
	plugName := c.Query("plugName")
	zipPath, err := a.autoCodePluginService.PubPlug(plugName)
	if err != nil {
		zap_logger.Error("打包失败!", zap.Error(err))
		response.FailWithMessage("打包失败"+err.Error(), c)
		return
	}
	response.OkWithMessage(fmt.Sprintf("打包成功,文件路径为:%s", zipPath), c)
}

// InitMenu
// @Tags      AutoCodePlugin
// @Summary   打包插件
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "打包插件成功"
// @Router    /autoCode/initMenu [post]
func (a *AutoCodePluginApi) InitMenu(c *gin.Context) {
	var menuInfo system.InitMenu
	err := c.ShouldBindJSON(&menuInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.autoCodePluginService.InitMenu(menuInfo)
	if err != nil {
		zap_logger.Error("创建初始化Menu失败!", zap.Error(err))
		response.FailWithMessage("创建初始化Menu失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("文件变更成功", c)
}

// InitAPI
// @Tags      AutoCodePlugin
// @Summary   打包插件
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "打包插件成功"
// @Router    /autoCode/initAPI [post]
func (a *AutoCodePluginApi) InitAPI(c *gin.Context) {
	var apiInfo system.InitApi
	err := c.ShouldBindJSON(&apiInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.autoCodePluginService.InitAPI(apiInfo)
	if err != nil {
		zap_logger.Error("创建初始化API失败!", zap.Error(err))
		response.FailWithMessage("创建初始化API失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("文件变更成功", c)
}