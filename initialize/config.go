package initialize

import (
	"fmt"
	"go-gin/global"
	"os"
	"path"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

func init() {
	v := viper.New()

	v.SetConfigType("yaml")
	v.SetConfigFile(defaultConfigFile)

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(fmt.Errorf("Fatal error unmarshal config: %s \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			panic(fmt.Errorf("Fatal error unmarshal config: %s \n", err))
		}
	})

	global.GVA_VP = v
	mkdir()
}

func mkdir() {
	// 创建 Runtime 目录
	os.MkdirAll(global.GVA_CONFIG.System.RuntimeRootPath, os.ModePerm)

	// 创建 Logs 目录
	os.MkdirAll(path.Join(
		global.GVA_CONFIG.System.RuntimeRootPath,
		global.GVA_CONFIG.Log.Dir),
		os.ModePerm)
}
