package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// config := config.GetConfig()

		// reqKey := c.Request.Header.Get("X-Auth-Key")
		// reqSecret := c.Request.Header.Get("X-Auth-Secret")

		// var key string
		// var secret string
		// if key = config.GetString("http.auth.key"); len(strings.TrimSpace(key)) == 0 {
		// 	c.AbortWithStatus(http.StatusInternalServerError)
		// }

		// if secret = config.GetString("http.auth.secret"); len(strings.TrimSpace(secret)) == 0 {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// }

		// if key != reqKey || secret != reqSecret {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }

		c.Next()
	}
}
