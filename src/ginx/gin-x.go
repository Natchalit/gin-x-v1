package ginx

import "github.com/gin-gonic/gin"

type Context struct {
	c *gin.Context
}

func (s *Context) GetContext() *gin.Context {
	if s == nil {
		return nil
	}
	return s.c
}
