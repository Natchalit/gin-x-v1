package mainx

import (
	"sync"
)

func Clean() {
	var wg sync.WaitGroup
	wg.Add(4)
	// go func() { defer wg.Done(); sqlx.Cleanups() }()
	wg.Wait()
}
