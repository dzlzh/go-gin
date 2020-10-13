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
	r := gin.New()
	r.Use(gin.Logger())
	global.GVA_LOG.Debug("use middleware Logger")
	r.Use(gin.Recovery())
	global.GVA_LOG.Debug("use middleware Recovery")
	if global.GVA_CONFIG.System.Tls {
		r.Use(middleware.Tls())
		global.GVA_LOG.Debug("use middleware Tls")
	}
	r.Use(middleware.Cors())
	global.GVA_LOG.Debug("use middleware Cors")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "ok",
			"data":    gin.H{},
		})
	})

	return r
}
