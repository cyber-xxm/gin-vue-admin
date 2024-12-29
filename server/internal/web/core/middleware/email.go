package middleware

import (
	"bytes"
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"
	tools "github.com/cyber-xxm/gin-vue-admin/internal/utils/plugin/email/utils"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"io"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorToEmail(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var username string
		claims, _ := utils.GetClaims(c)
		if claims.Username != "" {
			username = claims.Username
		} else {
			id, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
			user, err := userService.FindUserById(id)
			if err != nil {
				username = "Unknown"
			}
			username = user.Username
		}
		body, _ := io.ReadAll(c.Request.Body)
		// 再重新写回请求体body中，ioutil.ReadAll会清空c.Request.Body中的数据
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		record := system.SysOperationRecord{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
		}
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		status := c.Writer.Status()
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		str := "接收到的请求为" + record.Body + "\n" + "请求方式为" + record.Method + "\n" + "报错信息如下" + record.ErrorMessage + "\n" + "耗时" + latency.String() + "\n"
		if status != 200 {
			subject := username + "" + record.Ip + "调用了" + record.Path + "报错了"
			if err := tools.ErrorToEmail(subject, str); err != nil {
				zap_logger.Error("ErrorToEmail Failed, err:", zap.Error(err))
			}
		}
	}
}
