package ginx

import "github.com/gin-gonic/gin"

func Guard(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
