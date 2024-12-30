package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/response"
	zap_logger "github.com/cyber-xxm/gin-vue-admin/internal/utils/zap-logger"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewDictionaryApi(db *gorm.DB) *DictionaryApi {
	return &DictionaryApi{
		DictionaryService: service.NewDictionaryService(db),
	}
}

type DictionaryApi struct {
	DictionaryService *service.DictionaryService
}

// CreateSysDictionary
// @Tags      SysDictionary
// @Summary   创建SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionary           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{msg=string}  "创建SysDictionary"
// @Router    /sysDictionary/createSysDictionary [post]
func (a *DictionaryApi) CreateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.DictionaryService.CreateSysDictionary(dictionary)
	if err != nil {
		zap_logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysDictionary
// @Tags      SysDictionary
// @Summary   删除SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionary           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{msg=string}  "删除SysDictionary"
// @Router    /sysDictionary/deleteSysDictionary [delete]
func (a *DictionaryApi) DeleteSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.DictionaryService.DeleteSysDictionary(dictionary)
	if err != nil {
		zap_logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateSysDictionary
// @Tags      SysDictionary
// @Summary   更新SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysDictionary           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{msg=string}  "更新SysDictionary"
// @Router    /sysDictionary/updateSysDictionary [put]
func (a *DictionaryApi) UpdateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = a.DictionaryService.UpdateSysDictionary(&dictionary)
	if err != nil {
		zap_logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSysDictionary
// @Tags      SysDictionary
// @Summary   用id查询SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     system.SysDictionary                                       true  "ID或字典英名"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "用id查询SysDictionary"
// @Router    /sysDictionary/findSysDictionary [get]
func (a *DictionaryApi) FindSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindQuery(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysDictionary, err := a.DictionaryService.GetSysDictionary(dictionary.Type, dictionary.ID, dictionary.Status)
	if err != nil {
		zap_logger.Error("字典未创建或未开启!", zap.Error(err))
		response.FailWithMessage("字典未创建或未开启", c)
		return
	}
	response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "查询成功", c)
}

// GetSysDictionaryList
// @Tags      SysDictionary
// @Summary   分页获取SysDictionary列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取SysDictionary列表,返回包括列表,总数,页码,每页数量"
// @Router    /sysDictionary/getSysDictionaryList [get]
func (a *DictionaryApi) GetSysDictionaryList(c *gin.Context) {
	list, err := a.DictionaryService.GetSysDictionaryInfoList()
	if err != nil {
		zap_logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(list, "获取成功", c)
}
