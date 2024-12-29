package api

import "github.com/cyber-xxm/gin-vue-admin/internal/utils/plugin/announcement/service"

var (
	Api         = new(api)
	serviceInfo = service.Service.Info
)

type api struct{ Info info }
