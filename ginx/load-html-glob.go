package ginx

import (
	"path/filepath"

	"github.com/Natchalit/gin-x-v1/tox"
	"github.com/Natchalit/gin-x-v1/validx"
)

func (rx *RGX) LoadHTMLGlob(templatesFolder, relativePath, root *string) {

	if validx.IsEmptyPtr(templatesFolder) {
		templatesFolder = tox.StringPtr(filepath.Join("src/templates", "**", "*.html"))
	}

	if validx.IsEmptyPtr(relativePath) {
		relativePath = tox.StringPtr(`/css`)
	}

	if validx.IsEmptyPtr(root) {
		root = tox.StringPtr(`src/templates/css`)
	}

	rx.Router.Static(tox.String(relativePath), tox.String(root))

	rx.Router.LoadHTMLGlob(tox.String(templatesFolder))
}
