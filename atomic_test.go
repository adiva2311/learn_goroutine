package learn_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var counter int64 = 0
	var group = sync.WaitGroup{}

	for x := 1; x <= 1000; x++ {
		group.Add(1)
		go func() {
			for y := 1; y <= 100; y++ {
				atomic.AddInt64(&counter, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter =", counter)
}
