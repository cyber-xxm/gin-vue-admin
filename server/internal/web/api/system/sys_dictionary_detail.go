package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	system2 "github.com/cyber-xxm/gin-vue-admin/internal/models/request/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/response"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"
	zap_logger "github.com/cyber-xxm/gin-vue-admin/internal/utils/zap-logger"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewDictionaryDetailApi(db *gorm.DB) *DictionaryDetailApi {
	return &DictionaryDetailApi{
		DictionaryDetailService: service.NewDictionaryDetailService(db),
	}
}

type DictionaryDetailApi struct {
	DictionaryDetailService *service.DictionaryDetailService
}

// CreateSysDictionaryDetail
// @Tags      SysDictionaryDetail
// @Summary   创建SysDictionaryDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionaryDetail     true  "SysDictionaryDetail模型"
// @Success   200   {object}  response.Response{msg=string}  "创建SysDictionaryDetail"
// @Router    /sysDictionaryDetail/createSysDictionaryDetail [post]
func (a *DictionaryDetailApi) CreateSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.DictionaryDetailService.CreateSysDictionaryDetail(detail)
	if err != nil {
		zap_logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysDictionaryDetail
// @Tags      SysDictionaryDetail
// @Summary   删除SysDictionaryDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionaryDetail     true  "SysDictionaryDetail模型"
// @Success   200   {object}  response.Response{msg=string}  "删除SysDictionaryDetail"
// @Router    /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
func (a *DictionaryDetailApi) DeleteSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.DictionaryDetailService.DeleteSysDictionaryDetail(detail)
	if err != nil {
		zap_logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSysDictionaryDetail
// @Tags      SysDictionaryDetail
// @Summary   更新SysDictionaryDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionaryDetail     true  "更新SysDictionaryDetail"
// @Success   200   {object}  response.Response{msg=string}  "更新SysDictionaryDetail"
// @Router    /sysDictionaryDetail/updateSysDictionaryDetail [put]
func (a *DictionaryDetailApi) UpdateSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.DictionaryDetailService.UpdateSysDictionaryDetail(&detail)
	if err != nil {
		zap_logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysDictionaryDetail
// @Tags      SysDictionaryDetail
// @Summary   用id查询SysDictionaryDetail
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysDictionaryDetail                                 true  "用id查询SysDictionaryDetail"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "用id查询SysDictionaryDetail"
// @Router    /sysDictionaryDetail/findSysDictionaryDetail [get]
func (a *DictionaryDetailApi) FindSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	err := c.ShouldBindQuery(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(detail, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reSysDictionaryDetail, err := a.DictionaryDetailService.GetSysDictionaryDetail(detail.ID)
	if err != nil {
		zap_logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"reSysDictionaryDetail": reSysDictionaryDetail}, "查询成功", c)
}

// GetSysDictionaryDetailList
// @Tags      SysDictionaryDetail
// @Summary   分页获取SysDictionaryDetail列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.SysDictionaryDetailSearch                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取SysDictionaryDetail列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysDictionaryDetail/getSysDictionaryDetailList [get]
func (a *DictionaryDetailApi) GetSysDictionaryDetailList(c *gin.Context) {
	var pageInfo system2.SysDictionaryDetailSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := a.DictionaryDetailService.GetSysDictionaryDetailInfoList(pageInfo)
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
