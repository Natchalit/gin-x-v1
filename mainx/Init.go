package mainx

import (
	"log"
	"time"

	"github.com/Natchalit/gin-x-v1/logx"
	"github.com/joho/godotenv"
)

// init ทั้งหมด
func Init() (*InitType, error) {
	logx.Init()
	_ = godotenv.Load(".env")
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
