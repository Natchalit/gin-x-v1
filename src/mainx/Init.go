package mainx

import (
	"log"
	"time"

	logx "github.com/Natchalit/gin-x/src/log-x"
)

// init ทั้งหมด
func Init() (*InitType, error) {
	logx.Init()

	// Time Zone
	loc, ex := time.LoadLocation("Asia/Bangkok")
	if ex != nil {
		log.Fatalf("set time zone, %v", ex.Error())
	}
	time.Local = loc

	return &InitType{
		Location: loc,
	}, nil
}
