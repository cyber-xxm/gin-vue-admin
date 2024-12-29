package router

import "github.com/cyber-xxm/gin-vue-admin/internal/utils/plugin/announcement/api"

var (
	Router  = new(router)
	apiInfo = api.Api.Info
)

type router struct{ Info Info }
