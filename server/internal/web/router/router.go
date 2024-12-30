package router

import (
	"context"
	"github.com/cyber-xxm/gin-vue-admin/internal/initialize"
	models "github.com/cyber-xxm/gin-vue-admin/internal/models/config"
	zap_logger "github.com/cyber-xxm/gin-vue-admin/internal/utils/zap-logger"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/api/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/middleware"
	"github.com/cyber-xxm/gin-vue-admin/internal/web/core/plugin"
	systemRouter "github.com/cyber-xxm/gin-vue-admin/internal/web/router/system"
	service "github.com/cyber-xxm/gin-vue-admin/internal/web/service/system"
	"gorm.io/gorm"
	"net/http"
	"os"

	"github.com/cyber-xxm/gin-vue-admin/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 创建一个中间件来传递 context.Context
func contextMiddleware(parentCtx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建一个新的子上下文
		ctx, cancel := context.WithCancel(parentCtx)
		defer cancel() // 在请求结束时取消上下文

		// 将新的上下文存储在gin上下文中
		c.Set("rootCtx", ctx)

		// 继续处理请求
		c.Next()
	}
}

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

// Routers 初始化总路由
func Routers(rootCtx context.Context) *gin.Engine {
	cfg := rootCtx.Value("config").(*models.Server)
	db := rootCtx.Value("db").(*gorm.DB)
	// 初始化Gin
	r := gin.Default()

	// 使用中间件来传递上下文
	r.Use(contextMiddleware(rootCtx))
	r.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		r.Use(gin.Logger())
	}
	codeHistoryApi := system.NewAutoCodeHistoryApi(db)
	codePackageApi := system.NewAutoCodePackageApi(db)
	codePluginApi := system.NewAutoCodePluginApi(db)
	codeTemplateApi := system.NewAutoCodeTemplateApi(db)
	systemApi := system.NewSystemApi(db)
	authorityApi := system.NewAuthorityApi(db)
	authorityBtnApi := system.NewAuthorityBtnApi(db)
	autoCodeApi := system.NewAutoCodeApi(db)
	casbinApi := system.NewCasbinApi(db)
	configApi := system.NewConfigApi(db)
	dictionaryApi := system.NewDictionaryApi(db)
	dictionaryDetailApi := system.NewDictionaryDetailApi(db)
	exportTemplateApi := system.NewSysExportTemplateApi(db)
	dbApi := system.NewDBApi(db)
	jwtApi := system.NewJwtApi(db)
	menuApi := system.NewMenuApi(db)
	operationRecordApi := system.NewOperationRecordApi(db)
	sysParamsApi := system.NewSysParamsApi(db)
	baseApi := system.NewBaseApi(db)

	systemApiRouter := systemRouter.NewSystemApiRouter(systemApi)
	authorityRouter := systemRouter.NewAuthorityRouter(authorityApi)
	authorityBtnRouter := systemRouter.NewAuthorityBtnRouter(authorityBtnApi)
	autoCodeRouter := systemRouter.NewAutoCodeRouter(autoCodeApi, codeTemplateApi, codePackageApi, codePluginApi)
	codeHistoryRouter := systemRouter.NewAutoCodeHistoryRouter(codeHistoryApi)
	baseRouter := systemRouter.NewBaseRouter(baseApi)
	casbinRouter := systemRouter.NewCasbinRouter(casbinApi)
	dictionaryRouter := systemRouter.NewDictionaryRouter(dictionaryApi)
	dictionaryDetailRouter := systemRouter.NewDictionaryDetailRouter(dictionaryDetailApi)
	exportTemplateRouter := systemRouter.NewSysExportTemplateRouter(exportTemplateApi)
	initDbRouter := systemRouter.NewInitRouter(dbApi)
	jwtRouter := systemRouter.NewJwtRouter(jwtApi)
	menuRouter := systemRouter.NewMenuRouter(menuApi)
	operationRecordRouter := systemRouter.NewOperationRecordRouter(operationRecordApi)
	sysParamsRouter := systemRouter.NewSysParamsRouter(sysParamsApi)
	configRouter := systemRouter.NewConfigRouter(configApi)
	userRouter := systemRouter.NewUserRouter(baseApi)

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面3行注释
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/assets", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	r.StaticFS(cfg.Local.StorePath, justFilesFilesystem{http.Dir(cfg.Local.StorePath)}) // Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	// zap_logger.Info("use middleware cors")
	docs.SwaggerInfo.BasePath = cfg.System.RouterPrefix
	r.GET(cfg.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	zap_logger.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := r.Group(cfg.System.RouterPrefix)
	PrivateGroup := r.Group(cfg.System.RouterPrefix)

	PrivateGroup.Use(middleware.JWTAuth(service.NewJwtService(db))).Use(middleware.CasbinHandler(service.NewCasbinService(db)))

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		baseRouter.InitBaseRouter(PublicGroup)   // 注册基础功能路由 不做鉴权
		initDbRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}

	{
		systemApiRouter.InitApiRouter(PrivateGroup, PublicGroup, operationRecordApi.OperationRecordService)           // 注册功能api路由
		jwtRouter.InitJwtRouter(PrivateGroup)                                                                         // jwt相关路由
		userRouter.InitUserRouter(PrivateGroup, operationRecordApi.OperationRecordService)                            // 注册用户路由
		menuRouter.InitMenuRouter(PrivateGroup, operationRecordApi.OperationRecordService)                            // 注册menu路由
		configRouter.InitSystemRouter(PrivateGroup, operationRecordApi.OperationRecordService)                        // system相关路由
		casbinRouter.InitCasbinRouter(PrivateGroup, operationRecordApi.OperationRecordService)                        // 权限相关路由
		autoCodeRouter.InitAutoCodeRouter(PrivateGroup, PublicGroup)                                                  // 创建自动化代码
		authorityRouter.InitAuthorityRouter(PrivateGroup, operationRecordApi.OperationRecordService)                  // 注册角色路由
		dictionaryRouter.InitSysDictionaryRouter(PrivateGroup, operationRecordApi.OperationRecordService)             // 字典管理
		codeHistoryRouter.InitAutoCodeHistoryRouter(PrivateGroup)                                                     // 自动化代码历史
		operationRecordRouter.InitSysOperationRecordRouter(PrivateGroup)                                              // 操作记录
		dictionaryDetailRouter.InitSysDictionaryDetailRouter(PrivateGroup, operationRecordApi.OperationRecordService) // 字典详情管理
		authorityBtnRouter.InitAuthorityBtnRouterRouter(PrivateGroup, operationRecordApi.OperationRecordService)      // 按钮权限管理
		exportTemplateRouter.InitSysExportTemplateRouter(PrivateGroup, operationRecordApi.OperationRecordService)     // 导出模板
		sysParamsRouter.InitSysParamsRouter(PrivateGroup, operationRecordApi.OperationRecordService)                  // 参数管理
		exampleRouter.InitCustomerRouter(PrivateGroup)                                                                // 客户路由
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)                                                   // 文件上传下载功能路由

	}

	//插件路由安装
	plugin.InstallPlugin(PrivateGroup, PublicGroup, r)

	// 注册业务路由
	initialize.initBizRouter(PrivateGroup, PublicGroup)

	zap_logger.Info("router register success")
	return r
}
