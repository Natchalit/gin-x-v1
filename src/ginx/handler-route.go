package ginx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"

	"github.com/Natchalit/gin-x/src/errorx"
	logx "github.com/Natchalit/gin-x/src/log-x"
	"github.com/gin-gonic/gin"
)

type HandleR struct {
	StatusCode     int
	ResponseHeader map[string]string
	ResponseBody   any
}

func (r *RGX) _Handler(c *gin.Context, onWorkerX func(*Context) (any, error)) (hr *HandleR) {

	// start := time.Now()
	hr = &HandleR{
		ResponseHeader: map[string]string{},
	}

	// request id
	// s._RID = uuid.NewString()

	// https://blog.flexicondev.com/read-go-http-request-body-multiple-times
	var jbody any
	// _ = c.ShouldBindJSON(&jbody)
	body, _ := io.ReadAll(c.Request.Body)
	if len(body) > 0 {
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	}
	_ = json.Unmarshal(body, &jbody)

	// cluster call
	var respError string
	defer func() {

		// statusCode
		// statusCode := hr.StatusCode

		// resp error
		if r := recover(); r != nil {
			respError = fmt.Sprintf(`PANIC:%s`, r)
			statusCode := http.StatusInternalServerError
			hr.StatusCode = statusCode
			hr.ResponseBody = errorx.ErrorChk(c, statusCode, respError)
			logx.Errorf(string(debug.Stack()))
		}

		// respError
		if respError != `` {
			_ = c.Error(errorx.Internal(respError))
		}

	}()

	// context
	ctx := &Context{c: c}

	// worker
	rx, ex := onWorkerX(ctx)
	if ex != nil {

		// isNilResponse := validx.IsNil(rx)
		// if !isNilResponse {
		// 	hr.ResponseBody = rx
		// }
		// if v, ok := ex.(*errorx.EX); ok {
		// 	if v != nil {
		// 		respError = v.Message
		// 		if isNilResponse {
		// 			hr.ResponseBody = v
		// 		}
		// 		if v.StatusCode == 0 {
		// 			hr.StatusCode = http.StatusInternalServerError
		// 		} else {
		// 			hr.StatusCode = v.StatusCode
		// 		}
		// 	}
		// } else {
		// 	respError = ex.Error()
		// 	if isNilResponse {
		// 		hr.ResponseBody = errorx.ErrorC(c, http.StatusInternalServerError, ex.Error())
		// 	}
		// 	if hr.StatusCode == 0 {
		// 		hr.StatusCode = http.StatusInternalServerError
		// 	} else {
		// 		hr.StatusCode = v.StatusCode
		// 	}
		// }

	} else if rx != nil {
		if rc, ok := rx.(*_RC); ok {
			hr.StatusCode = rc.StatusCode
			hr.ResponseBody = rc.ResponseBody
		} else {
			hr.StatusCode = http.StatusOK
			hr.ResponseBody = rx
		}
	}

	return
}

type _RC struct {
	StatusCode   int
	ResponseBody any
}
