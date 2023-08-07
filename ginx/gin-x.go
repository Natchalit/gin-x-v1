package ginx

import (
	"net/http"

	"github.com/Natchalit/gin-x-v1/errorx"
	"github.com/gin-gonic/gin"
)

type Context struct {
	c *gin.Context
}

func (s *Context) GetContext() *gin.Context {
	if s == nil {
		return nil
	}
	return s.c
}

func (s *Context) ShouldBindJSON(obj any) error {
	return s.GetContext().ShouldBindJSON(obj)
}

func (s *Context) Error(err error) error {
	if err == nil {
		return nil
	}

	if ex, ok := err.(*errorx.EX); ok {
		return ex
	}

	return errorx.ErrorChk(s.GetContext(), http.StatusInternalServerError, err.Error())
}
