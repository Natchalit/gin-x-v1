package ginx

import (
	"net/http"

	"github.com/Natchalit/gin-x-v1/errorx"
)

func (c *Context) BadRequest(f string, a ...any) error {
	return errorx.ErrorChk(c.c, http.StatusBadRequest, f, a...)
}

func (c *Context) InternalServerError(f string, a ...any) error {
	return errorx.ErrorChk(c.c, http.StatusInternalServerError, f, a...)
}

func BadRequest(f string, a ...any) error {
	return errorx.ErrorChk(nil, http.StatusBadRequest, f, a...)
}

func InternalServerError(f string, a ...any) error {
	return errorx.ErrorChk(nil, http.StatusInternalServerError, f, a...)
}
