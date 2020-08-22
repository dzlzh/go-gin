package main

import (
	"go-gin/global"
	"go-gin/initialize"
)

func main() {
	// 初始化 Mysql
	initialize.Mysql()
	defer global.GVA_DB.Close()
	// 初始化 Redis
	initialize.Redis()
}
