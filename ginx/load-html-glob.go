package ginx

import (
	"path/filepath"

	"github.com/Natchalit/gin-x-v1/validx"
)

func (rx *RGX) LoadHTMLGlob() {

	templatesFolder := rx.TemplatesFolder
	relativePath := rx.RelativePath
	root := rx.Root

	if validx.IsEmpty(templatesFolder) {
		templatesFolder = filepath.Join("src/templates", "**", "*.html")
	}

	if validx.IsEmpty(relativePath) {
		relativePath = `/css`
	}

	if validx.IsEmpty(root) {
		root = `src/templates/css`
	}

	rx.Router.Static(relativePath, root)

	rx.Router.LoadHTMLGlob(templatesFolder)
}
