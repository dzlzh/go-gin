package bootstrap

import (
	"fmt"
	"go-gin/global"
	"go-gin/initialize"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func Run() {
	// 加载路由
	router := initialize.Routers()

	// 启动服务
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, router)
	global.GVA_LOG.Debug("server run success on ", address)
	global.GVA_LOG.Error(s.ListenAndServe())
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
