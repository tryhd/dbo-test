package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tryhd/dbo-test/app/config"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := config.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
