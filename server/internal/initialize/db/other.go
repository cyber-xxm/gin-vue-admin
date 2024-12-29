package db

import (
	"bufio"
	"context"
	models "github.com/cyber-xxm/gin-vue-admin/internal/models/config"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"os"
	"strings"
)

func OtherInit(ctx context.Context) local_cache.Cache {
	cfg := ctx.Value("config").(models.Server)
	dr, err := utils.ParseDuration(cfg.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(cfg.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	cache := local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
	file, err := os.Open("go.mod")
	if err == nil && cfg.AutoCode.Module == "" {
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		cfg.AutoCode.Module = strings.TrimPrefix(scanner.Text(), "module ")
	}
	return cache
}
