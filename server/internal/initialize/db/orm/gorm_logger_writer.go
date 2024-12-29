package orm

import (
	"fmt"
	"gorm.io/gorm/logger"
)

type Writer struct {
	config GeneralDB
	writer logger.Writer
}

func NewWriter(config GeneralDB) *Writer {
	return &Writer{config: config}
}

// Printf 格式化打印日志
func (c *Writer) Printf(message string, data ...any) {
	if c.config.LogZap {
		switch c.config.LogLevel() {
		case logger.Silent:
			fmt.Println("silent", fmt.Sprintf(message, data...))
		case logger.Error:
			fmt.Println("error", fmt.Sprintf(message, data...))
		case logger.Warn:
			fmt.Println("warn", fmt.Sprintf(message, data...))
		case logger.Info:
			fmt.Println("info", fmt.Sprintf(message, data...))
		default:
			fmt.Println(fmt.Sprintf(message, data...))
		}
		return
	}
}
