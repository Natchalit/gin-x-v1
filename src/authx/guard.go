package authx

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func Guard(c *gin.Context) error {
	s := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")

	if ex := ValidEmptyToken(token); ex != nil {
		return ex
	}

	return nil
}
