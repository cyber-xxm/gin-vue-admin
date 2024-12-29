package zap_logger

import (
	"fmt"
	"github.com/cyber-xxm/gin-vue-admin/internal/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger

func NewZapLogger(l *zap.Logger) {
	logger = l
}

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap(dir string, levels []zapcore.Level, showLine bool) (logger *zap.Logger) {
	if ok, _ := utils.PathExists(dir); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", dir)
		_ = os.Mkdir(dir, os.ModePerm)
	}
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if showLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func Error(s string, f ...zap.Field) {
	logger.Error(s, f...)
}

func Info(s string, f ...zap.Field) {
	logger.Info(s, f...)
}

func Debug(s string, f ...zap.Field) {
	logger.Debug(s, f...)
}

func Warn(s string, f ...zap.Field) {
	logger.Warn(s, f...)
}
