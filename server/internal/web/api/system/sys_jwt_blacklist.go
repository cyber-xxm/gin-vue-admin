package system

import (
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/response"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewJwtApi(db *gorm.DB) *JwtApi {
	return &JwtApi{
		jwtService: service.NewJwtService(db),
	}
}

type JwtApi struct {
	jwtService *service.JwtService
}

// JsonInBlacklist
// @Tags      Jwt
// @Summary   jwt加入黑名单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "jwt加入黑名单"
// @Router    /jwt/jsonInBlacklist [post]
func (a *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := utils.GetToken(c)
	jwt := system.JwtBlacklist{Jwt: token}
	err := a.jwtService.JsonInBlacklist(jwt)
	if err != nil {
		zap_logger.Error("jwt作废失败!", zap.Error(err))
		response.FailWithMessage("jwt作废失败", c)
		return
	}
	utils.ClearToken(c)
	response.OkWithMessage("jwt作废成功", c)
}
