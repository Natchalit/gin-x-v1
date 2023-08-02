package ginx

import "github.com/gin-gonic/gin"

func UseHeader(c *gin.Context) {
	c.Header("Content-Type", "image/jpeg")
}
