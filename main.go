package main

import (
	"go-gin/bootstrap"
	"go-gin/global"
	"go-gin/initialize"
)

func main() {
	// 初始化 Log
	initialize.Log()
	// 初始化 Mysql
	initialize.Mysql()
	defer func() {
		db, _ := global.GVA_DB.DB()
		db.Close()
	}()
	// 初始化 Redis
	initialize.Redis()

	bootstrap.Run()
}
