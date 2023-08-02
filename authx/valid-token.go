package authx

import (
	"net/http"

	"github.com/Natchalit/gin-x-v1/errorx"
	"github.com/Natchalit/gin-x-v1/validx"
)

func ValidEmptyToken(token string) error {

	if validx.IsEmpty(token) || token != `LOGIN_TOKEN` {
		massage := `Invalid Auth`
		return errorx.ErrorChk(nil, http.StatusUnauthorized, massage)
	}

	return nil
}
