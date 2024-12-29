package system

import (
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request"
	systemReq "github.com/cyber-xxm/gin-vue-admin/internal/models/request/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/response"
	systemRes "github.com/cyber-xxm/gin-vue-admin/internal/models/response/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewSystemApiApi(db *gorm.DB) *SystemApiApi {
	return &SystemApiApi{
		apiService: service.NewApiService(db),
	}
}

type SystemApiApi struct {
	apiService *service.ApiService
}

// CreateApi
// @Tags      SysApi
// @Summary   创建基础api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysApi                  true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "创建基础api"
// @Router    /api/createApi [post]
func (a *SystemApiApi) CreateApi(c *gin.Context) {
	var api system.SysApi
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.apiService.CreateApi(api)
	if err != nil {
		zap_logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// SyncApi
// @Tags      SysApi
// @Summary   同步API
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "同步API"
// @Router    /api/syncApi [get]
func (a *SystemApiApi) SyncApi(c *gin.Context) {
	newApis, deleteApis, ignoreApis, err := a.apiService.SyncApi()
	if err != nil {
		zap_logger.Error("同步失败!", zap.Error(err))
		response.FailWithMessage("同步失败", c)
		return
	}
	response.OkWithData(gin.H{
		"newApis":    newApis,
		"deleteApis": deleteApis,
		"ignoreApis": ignoreApis,
	}, c)
}

// GetApiGroups
// @Tags      SysApi
// @Summary   获取API分组
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "获取API分组"
// @Router    /api/getApiGroups [get]
func (a *SystemApiApi) GetApiGroups(c *gin.Context) {
	groups, apiGroupMap, err := a.apiService.GetApiGroups()
	if err != nil {
		zap_logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(gin.H{
		"groups":      groups,
		"apiGroupMap": apiGroupMap,
	}, c)
}

// IgnoreApi
// @Tags      IgnoreApi
// @Summary   忽略API
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "同步API"
// @Router    /api/ignoreApi [post]
func (a *SystemApiApi) IgnoreApi(c *gin.Context) {
	var ignoreApi system.SysIgnoreApi
	err := c.ShouldBindJSON(&ignoreApi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.apiService.IgnoreApi(ignoreApi)
	if err != nil {
		zap_logger.Error("忽略失败!", zap.Error(err))
		response.FailWithMessage("忽略失败", c)
		return
	}
	response.Ok(c)
}

// EnterSyncApi
// @Tags      SysApi
// @Summary   确认同步API
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "确认同步API"
// @Router    /api/enterSyncApi [post]
func (a *SystemApiApi) EnterSyncApi(c *gin.Context) {
	var syncApi systemRes.SysSyncApis
	err := c.ShouldBindJSON(&syncApi)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.apiService.EnterSyncApi(syncApi)
	if err != nil {
		zap_logger.Error("忽略失败!", zap.Error(err))
		response.FailWithMessage("忽略失败", c)
		return
	}
	response.Ok(c)
}

// DeleteApi
// @Tags      SysApi
// @Summary   删除api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysApi                  true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除api"
// @Router    /api/deleteApi [post]
func (a *SystemApiApi) DeleteApi(c *gin.Context) {
	var api system.SysApi
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.apiService.DeleteApi(api)
	if err != nil {
		zap_logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetApiList
// @Tags      SysApi
// @Summary   分页获取API列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SearchApiParams                               true  "分页获取API列表"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取API列表,返回包括列表,总数,页码,每页数量"
// @Router    /api/getApiList [post]
func (a *SystemApiApi) GetApiList(c *gin.Context) {
	var pageInfo systemReq.SearchApiParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := a.apiService.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		zap_logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetApiById
// @Tags      SysApi
// @Summary   根据id获取api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                                   true  "根据id获取api"
// @Success   200   {object}  response.Response{data=systemRes.SysAPIResponse}  "根据id获取api,返回包括api详情"
// @Router    /api/getApiById [post]
func (a *SystemApiApi) GetApiById(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	api, err := a.apiService.GetApiById(idInfo.ID)
	if err != nil {
		zap_logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysAPIResponse{Api: api}, "获取成功", c)
}

// UpdateApi
// @Tags      SysApi
// @Summary   修改基础api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysApi                  true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "修改基础api"
// @Router    /api/updateApi [post]
func (a *SystemApiApi) UpdateApi(c *gin.Context) {
	var api system.SysApi
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.apiService.UpdateApi(api)
	if err != nil {
		zap_logger.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// GetAllApis
// @Tags      SysApi
// @Summary   获取所有的Api 不分页
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysAPIListResponse,msg=string}  "获取所有的Api 不分页,返回包括api列表"
// @Router    /api/getAllApis [post]
func (a *SystemApiApi) GetAllApis(c *gin.Context) {
	authorityID := utils.GetUserAuthorityId(c)
	apis, err := a.apiService.GetAllApis(authorityID)
	if err != nil {
		zap_logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysAPIListResponse{Apis: apis}, "获取成功", c)
}

// DeleteApisByIds
// @Tags      SysApi
// @Summary   删除选中Api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除选中Api"
// @Router    /api/deleteApisByIds [delete]
func (a *SystemApiApi) DeleteApisByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.apiService.DeleteApisByIds(ids)
	if err != nil {
		zap_logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// FreshCasbin
// @Tags      SysApi
// @Summary   刷新casbin缓存
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "刷新成功"
// @Router    /api/freshCasbin [get]
func (a *SystemApiApi) FreshCasbin(c *gin.Context) {
	err := a.casbinService.FreshCasbin()
	if err != nil {
		zap_logger.Error("刷新失败!", zap.Error(err))
		response.FailWithMessage("刷新失败", c)
		return
	}
	response.OkWithMessage("刷新成功", c)
}
