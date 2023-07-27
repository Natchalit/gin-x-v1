package errorx

import (
	"fmt"
	"net/http"
)

func Error(err error) error {

	ex, ok := err.(*EX)
	if !ok {
		ex = &EX{
			Message: err.Error(),
		}
	}

	if ex.StatusCode == 0 {
		ex.StatusCode = http.StatusInternalServerError
	}

	if ex.Message == `` {
		ex.Message = `[EMPTY]`
	}

	return ex
}

func Internal(f string, a ...any) error {
	return Error(&EX{
		Message:    fmt.Sprintf(f, a...),
		StatusCode: http.StatusInternalServerError,
	})
}

func BadRequest(f string, a ...any) error {
	return Error(&EX{
		Message:    fmt.Sprintf(f, a...),
		StatusCode: http.StatusBadRequest,
	})
}
