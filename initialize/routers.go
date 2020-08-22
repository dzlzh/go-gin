package initialize

import (
	"go-gin/global"
	"go-gin/middleware"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	if global.GVA_CONFIG.System.Debug {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}
	r := gin.Default()
	global.GVA_LOG.Debug("use middleware Logger and Recovery")

	return r
}
