package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		uid := c.Request.Header.Get("x-uid")
		if token == "" || uid == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "无权限",
				"data": gin.H{},
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
