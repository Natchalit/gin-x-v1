package errorx

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorChk(c *gin.Context, statusCode int, f string, a ...any) error {

	// create new error
	ex := EX{
		Message: fmt.Sprintf(f, a...),
	}

	if c == nil {
		ex.StatusCode = c.Writer.Status()
		ex.ClientIP = c.ClientIP()
		ex.Method = c.Request.Method
		ex.URL = c.Request.Host + c.Request.URL.Path
	}
	if statusCode > 0 {
		ex.StatusCode = statusCode
	}

	return Error(&ex)
}
