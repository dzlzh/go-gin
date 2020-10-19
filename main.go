package main

import (
	"go-gin/bootstrap"
	"go-gin/initialize"
)

func main() {
	// 初始化 Log
	initialize.Log()
	// 初始化 Mysql
	initialize.Mysql()
	defer initialize.MysqlClose()
	// 初始化 Redis
	initialize.Redis()

	bootstrap.Run()
}
