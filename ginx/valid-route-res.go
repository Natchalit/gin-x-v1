package ginx

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (r *RGX) _ValidRouteRes(onWorkerX func(*Context) (any, error)) []gin.HandlerFunc {

	handlers := []gin.HandlerFunc{}
	handlers = append(handlers, func(ctx *gin.Context) {
		hr := r._Handler(ctx, onWorkerX)
		for k, v := range hr.ResponseHeader {
			ctx.Header(k, v)
		}
		if hr.StatusCode == http.StatusNoContent {
			ctx.AbortWithStatus(http.StatusNoContent)
		} else {
			if hr.ResponseBody == nil {
				//c.AbortWithStatus(http.StatusNoContent)
			} else if vx, ok := hr.ResponseBody.(string); ok {
				if vs := strings.TrimSpace(vx); vs == `` || vs == `{}` {
					// c.AbortWithStatus(http.StatusNoContent)
				} else {
					if json.Valid([]byte(vx)) {
						res := hr.ResponseBody.(map[string]interface{})
						if res[`isWeb`].(bool) {
							path := res[`path`].(*string)
							ctx.HTML(hr.StatusCode, *path, res[`data`])
						} else {
							ctx.JSON(hr.StatusCode, res[`data`])
						}
					} else {
						ctx.String(hr.StatusCode, vx)
					}
				}
			} else {
				if res, ok := hr.ResponseBody.(map[string]interface{}); ok {
					if res[`isWeb`].(bool) {
						path := res[`path`].(*string)
						ctx.HTML(hr.StatusCode, *path, res[`data`])
					} else {
						ctx.JSON(hr.StatusCode, res[`data`])
					}
				} else {
					ctx.JSON(hr.StatusCode, hr.ResponseBody)
				}
			}
		}
	})

	return handlers
}
