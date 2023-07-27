package authx

import (
	"net/http"

	"github.com/Natchalit/gin-x/errorx"
	"github.com/Natchalit/gin-x/validx"
)

func ValidEmptyToken(token string) error {

	if validx.IsEmpty(token) || token != `LOGIN_TOKEN` {
		massage := `Invalid Auth`
		return errorx.ErrorChk(nil, http.StatusUnauthorized, massage)
	}

	return nil
}
