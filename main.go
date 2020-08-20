package main

import (
	"go-gin/global"
	"go-gin/initialize"
)

func main() {
	initialize.Mysql()
	defer global.GVA_DB.Close()
}
