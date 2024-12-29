package initialize

import (
	"context"
	"github.com/cyber-xxm/gin-vue-admin/global"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils/plugin/announcement/router"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(engine *gin.Engine, ctx context.Context) {
	db := ctx.Value("db").(*gorm.DB)
	jwtService := service.NewJwtService(db)
	casbinService := service.NewCasbinService(db)
	public := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private.Use(middleware.JWTAuth(jwtService)).Use(middleware.CasbinHandler(casbinService))
	router.Router.Info.Init(public, private)
}
