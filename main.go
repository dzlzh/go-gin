package main

import (
	"go-gin/cmd"
	"go-gin/initialize"
	"os"
)

func main() {
	// 初始化 Log
	initialize.Log()
	// 初始化 Mysql
	initialize.Mysql()
	defer initialize.MysqlClose()
	// 初始化 Redis
	initialize.Redis()

	err := cmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}
