package router

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/utils/plugin/email/api"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewEmailRouter(db *gorm.DB) *EmailRouter {
	return &EmailRouter{
		operationRecordService: service.NewOperationRecordService(db),
	}
}

type EmailRouter struct {
	operationRecordService *service.OperationRecordService
}

func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Use(middleware.OperationRecord(s.operationRecordService))
	EmailApi := api.ApiGroupApp.EmailApi.EmailTest
	SendEmail := api.ApiGroupApp.EmailApi.SendEmail
	{
		emailRouter.POST("emailTest", EmailApi)  // 发送测试邮件
		emailRouter.POST("sendEmail", SendEmail) // 发送邮件
	}
}
