package system

import (
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/response"
	zap_logger "github.com/cyber-xxm/gin-vue-admin/internal/utils/zap-logger"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func NewDBApi(db *gorm.DB) *DBApi {
	return &DBApi{
		InitDBService: service.NewInitDBService(db),
	}
}

type DBApi struct {
	InitDBService *service.InitDBService
}

// InitDB
// @Tags     InitDB
// @Summary  初始化用户数据库
// @Produce  application/json
// @Param    data  body      request.InitDB                  true  "初始化数据库参数"
// @Success  200   {object}  response.Response{data=string}  "初始化用户数据库"
// @Router   /init/initdb [post]
func (a *DBApi) InitDB(c *gin.Context) {
	if global.GVA_DB != nil {
		zap_logger.Error("已存在数据库配置!")
		response.FailWithMessage("已存在数据库配置", c)
		return
	}
	var dbInfo system.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		zap_logger.Error("参数校验不通过!", zap.Error(err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	if err := a.InitDBService.InitDB(dbInfo); err != nil {
		zap_logger.Error("自动创建数据库失败!", zap.Error(err))
		response.FailWithMessage("自动创建数据库失败，请查看后台日志，检查后在进行初始化", c)
		return
	}
	response.OkWithMessage("自动创建数据库成功", c)
}

// CheckDB
// @Tags     CheckDB
// @Summary  初始化用户数据库
// @Produce  application/json
// @Success  200  {object}  response.Response{data=map[string]interface{},msg=string}  "初始化用户数据库"
// @Router   /init/checkdb [post]
func (a *DBApi) CheckDB(c *gin.Context) {
	var (
		message  = "前往初始化数据库"
		needInit = true
	)

	if global.GVA_DB != nil {
		message = "数据库无需初始化"
		needInit = false
	}
	zap_logger.Info(message)
	response.OkWithDetailed(gin.H{"needInit": needInit}, message, c)
}
