package system

import (
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/response"
	systemRes "github.com/cyber-xxm/gin-vue-admin/internal/models/response/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewSystemConfigApi(db *gorm.DB) *SystemConfigApi {
	return &SystemConfigApi{
		systemConfigService: service.NewSystemConfigService(db),
	}
}

type SystemConfigApi struct {
	systemConfigService *service.SystemConfigService
}

// GetSystemConfig
// @Tags      System
// @Summary   获取配置文件内容
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysConfigResponse,msg=string}  "获取配置文件内容,返回包括系统配置"
// @Router    /system/getSystemConfig [post]
func (a *SystemConfigApi) GetSystemConfig(c *gin.Context) {
	config, err := a.systemConfigService.GetSystemConfig()
	if err != nil {
		zap_logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "获取成功", c)
}

// SetSystemConfig
// @Tags      System
// @Summary   设置配置文件内容
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      system.System                   true  "设置配置文件内容"
// @Success   200   {object}  response.Response{data=string}  "设置配置文件内容"
// @Router    /system/setSystemConfig [post]
func (a *SystemConfigApi) SetSystemConfig(c *gin.Context) {
	var sys system.System
	err := c.ShouldBindJSON(&sys)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.systemConfigService.SetSystemConfig(sys)
	if err != nil {
		zap_logger.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// ReloadSystem
// @Tags      System
// @Summary   重启系统
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "重启系统"
// @Router    /system/reloadSystem [post]
func (a *SystemConfigApi) ReloadSystem(c *gin.Context) {
	err := utils.Reload()
	if err != nil {
		zap_logger.Error("重启系统失败!", zap.Error(err))
		response.FailWithMessage("重启系统失败", c)
		return
	}
	response.OkWithMessage("重启系统成功", c)
}

// GetServerInfo
// @Tags      System
// @Summary   获取服务器信息
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取服务器信息"
// @Router    /system/getServerInfo [post]
func (a *SystemConfigApi) GetServerInfo(c *gin.Context) {
	server, err := a.systemConfigService.GetServerInfo()
	if err != nil {
		zap_logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
}
