package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Options(allow []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.Header("Allow", strings.Join(allow, ","))
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
