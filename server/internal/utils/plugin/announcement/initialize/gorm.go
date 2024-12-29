package initialize

import (
	"context"
	"fmt"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils/plugin/announcement/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm(ctx context.Context) {
	db := ctx.Value("db").(*gorm.DB)
	err := db.WithContext(ctx).AutoMigrate(
		new(model.Info),
	)
	if err != nil {
		err = errors.Wrap(err, "注册表失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
}
