package config

import (
	"flag"
	"fmt"
	models "github.com/cyber-xxm/gin-vue-admin/internal/models/config"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/cyber-xxm/gin-vue-admin/global"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper(path ...string) (*viper.Viper, models.Server) {
	var conf string

	if len(path) == 0 {
		flag.StringVar(&conf, "c", "", "choose config file.")
		flag.Parse()
		if conf == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				switch gin.Mode() {
				case gin.DebugMode:
					conf = global.ConfigDefaultFile
				case gin.ReleaseMode:
					conf = global.ConfigReleaseFile
				case gin.TestMode:
					conf = global.ConfigTestFile
				}
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.Mode(), conf)
			} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				conf = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", global.ConfigEnv, conf)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", conf)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		conf = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", conf)
	}

	v := viper.New()
	v.SetConfigFile(conf)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	var cfg models.Server
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&cfg); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	cfg.AutoCode.Root, _ = filepath.Abs("..")
	return v, cfg
}
