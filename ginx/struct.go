package ginx

import "github.com/gin-gonic/gin"

type RTX struct {
	PORT   int
	MODULE string
}

type RGX struct {
	SRV             string
	Router          *gin.Engine
	TemplatesFolder string
	RelativePath    string
	Root            string
}

func (rx *RGX) _HandlerX(g *gin.RouterGroup, relativePath string, onWorkerX func(*Context) (any, error), handlers ...gin.HandlerFunc) []gin.HandlerFunc {

	// set path
	path := getPath(g.BasePath(), relativePath)

	handlerx := []gin.HandlerFunc{}
	handlerx = append(handlerx, rx._ValidRouteRes(onWorkerX)...)
	handlerx = append(handlerx, Guard(path))
	handlerx = append(handlerx, handlers...)

	return handlerx

}
