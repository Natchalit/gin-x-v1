package authx

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Guard() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := c.Request.Header.Get("Authorization")
		token := strings.TrimPrefix(s, "Bearer ")

		if ex := ValidEmptyToken(token); ex != nil {
			c.JSON(http.StatusUnauthorized, "Invalid Token(41)")
			c.Abort()
			return
		}

		c.Next()
	}
}
