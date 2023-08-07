package gormx

import (
	"fmt"
	"strings"

	"github.com/Natchalit/gin-x-v1/ginx"
)

func XAutoMigrate(c *ginx.Context) (any, error) {

	type DBName struct {
		DBName string `json:"db_name"`
	}

	Found := false

	dto := DBName{}
	if ex := c.ShouldBindJSON(&dto); ex != nil {
		return nil, ex
	}

	for k, v := range Migration {
		if strings.EqualFold(k, dto.DBName) {
			v(k)
			Found = true
		}
	}

	if !Found {
		return nil, c.Error(fmt.Errorf(`not found database`))
	}

	return nil, nil
}
