package system

import (
	"fmt"
	"github.com/cyber-xxm/gin-vue-admin/internal/models"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/response"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils/request"
	zap_logger "github.com/cyber-xxm/gin-vue-admin/internal/utils/zap-logger"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"io"
	"strings"

	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewAutoCodeApi(db *gorm.DB) *AutoCodeApi {
	return &AutoCodeApi{
		AutoCodeService: service.NewAutoCodeHistoryService(db),
	}
}

type AutoCodeApi struct {
	AutoCodeService interface{}
}

// GetDB
// @Tags      AutoCode
// @Summary   获取当前所有数据库
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取当前所有数据库"
// @Router    /autoCode/getDB [get]
func (a *AutoCodeApi) GetDB(c *gin.Context) {
	businessDB := c.Query("businessDB")
	dbs, err := a.AutoCodeService.Database(businessDB).GetDB(businessDB)
	var dbList []map[string]interface{}
	for _, db := range global.GVA_CONFIG.DBList {
		var item = make(map[string]interface{})
		item["aliasName"] = db.AliasName
		item["dbName"] = db.Dbname
		item["disable"] = db.Disable
		item["dbtype"] = db.Type
		dbList = append(dbList, item)
	}
	if err != nil {
		zap_logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"dbs": dbs, "dbList": dbList}, "获取成功", c)
	}
}

// GetTables
// @Tags      AutoCode
// @Summary   获取当前数据库所有表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取当前数据库所有表"
// @Router    /autoCode/getTables [get]
func (a *AutoCodeApi) GetTables(c *gin.Context) {
	dbName := c.Query("dbName")
	businessDB := c.Query("businessDB")
	if dbName == "" {
		dbName = *global.GVA_ACTIVE_DBNAME
		if businessDB != "" {
			for _, db := range global.GVA_CONFIG.DBList {
				if db.AliasName == businessDB {
					dbName = db.Dbname
				}
			}
		}
	}

	tables, err := a.AutoCodeService.Database(businessDB).GetTables(businessDB, dbName)
	if err != nil {
		zap_logger.Error("查询table失败!", zap.Error(err))
		response.FailWithMessage("查询table失败", c)
	} else {
		response.OkWithDetailed(gin.H{"tables": tables}, "获取成功", c)
	}
}

// GetColumn
// @Tags      AutoCode
// @Summary   获取当前表所有字段
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取当前表所有字段"
// @Router    /autoCode/getColumn [get]
func (a *AutoCodeApi) GetColumn(c *gin.Context) {
	businessDB := c.Query("businessDB")
	dbName := c.Query("dbName")
	if dbName == "" {
		dbName = *global.GVA_ACTIVE_DBNAME
		if businessDB != "" {
			for _, db := range global.GVA_CONFIG.DBList {
				if db.AliasName == businessDB {
					dbName = db.Dbname
				}
			}
		}
	}
	tableName := c.Query("tableName")
	columns, err := a.AutoCodeService.Database(businessDB).GetColumn(businessDB, tableName, dbName)
	if err != nil {
		zap_logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"columns": columns}, "获取成功", c)
	}
}

func (a *AutoCodeApi) LLMAuto(c *gin.Context) {
	var llm models.JSONMap
	err := c.ShouldBindJSON(&llm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if global.GVA_CONFIG.AutoCode.AiPath == "" {
		response.FailWithMessage("请先前往插件市场个人中心获取AiPath并填入config.yaml中", c)
		return
	}

	path := strings.ReplaceAll(global.GVA_CONFIG.AutoCode.AiPath, "{FUNC}", fmt.Sprintf("api/chat/%s", llm["mode"]))
	res, err := request.HttpRequest(
		path,
		"POST",
		nil,
		nil,
		llm,
	)
	if err != nil {
		zap_logger.Error("大模型生成失败!", zap.Error(err))
		response.FailWithMessage("大模型生成失败"+err.Error(), c)
		return
	}
	var resStruct response.Response
	b, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		zap_logger.Error("大模型生成失败!", zap.Error(err))
		response.FailWithMessage("大模型生成失败"+err.Error(), c)
		return
	}
	err = json.Unmarshal(b, &resStruct)
	if err != nil {
		zap_logger.Error("大模型生成失败!", zap.Error(err))
		response.FailWithMessage("大模型生成失败"+err.Error(), c)
		return
	}

	if resStruct.Code == 7 {
		zap_logger.Error("大模型生成失败!"+resStruct.Msg, zap.Error(err))
		response.FailWithMessage("大模型生成失败"+resStruct.Msg, c)
		return
	}
	response.OkWithData(resStruct.Data, c)
}
