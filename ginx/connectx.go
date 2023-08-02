package ginx

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (rgx *RTX) Connect(routerx func(*RGX, *gin.RouterGroup)) {
	rx := RGX{
		TemplatesFolder: "",
		RelativePath:    "",
		Root:            "",
	}
	// c := gin.Default()
	gin.ForceConsoleColor()
	router := gin.New()

	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/"},
	}))

	router.Use(gin.Recovery())

	rx.Router = router
	rx.SRV = rgx.MODULE

	routerx(&rx, router.Group(rgx.MODULE))

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// _ = router.Run(fmt.Sprintf(":%v", rgx.PORT))

	s.ListenAndServe()
}
