package system

import (
	"github.com/cyber-xxm/gin-vue-admin/internal/models/db/system"
	"github.com/cyber-xxm/gin-vue-admin/internal/models/request"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
