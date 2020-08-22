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
	if global.GVA_CONFIG.System.Tls {
		r.Use(middleware.Tls())
		global.GVA_LOG.Debug("use middleware Tls")
	}
	r.Use(middleware.Cors())
	global.GVA_LOG.Debug("use middleware Cors")

	return r
}
