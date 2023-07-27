package ginx

import "github.com/gin-gonic/gin"

func (rx *RGX) GET(g *gin.RouterGroup, relativePath string, onWorkerX func(*Context) (any, error), handlers ...gin.HandlerFunc) {

	handlerx := rx._HandlerX(g, relativePath, onWorkerX, handlers...)

	g.GET(relativePath, handlerx...)
}

func (rx *RGX) POST(g *gin.RouterGroup, relativePath string, onWorkerX func(*Context) (any, error), handlers ...gin.HandlerFunc) {

	handlerx := rx._HandlerX(g, relativePath, onWorkerX, handlers...)

	g.POST(relativePath, handlerx...)
}

func (rx *RGX) PUT(g *gin.RouterGroup, relativePath string, onWorkerX func(*Context) (any, error), handlers ...gin.HandlerFunc) {

	handlerx := rx._HandlerX(g, relativePath, onWorkerX, handlers...)

	g.PUT(relativePath, handlerx...)
}

func (rx *RGX) DELETE(g *gin.RouterGroup, relativePath string, onWorkerX func(*Context) (any, error), handlers ...gin.HandlerFunc) {

	handlerx := rx._HandlerX(g, relativePath, onWorkerX, handlers...)

	g.DELETE(relativePath, handlerx...)
}
