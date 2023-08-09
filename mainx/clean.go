package mainx

import (
	"sync"

	"github.com/Natchalit/gin-x-v1/sqlx"
)

func Clean() {
	var wg sync.WaitGroup
	wg.Add(4)
	go func() { defer wg.Done(); sqlx.CleanConnections() }()
	wg.Wait()
}
