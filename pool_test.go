package learn_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool = sync.Pool{
		// Mengatur nilai default dari pool yang diberikan sehingga data tidak <nill>
		New: func() any {
			return "New User"
		},
	}

	pool.Put("Adiva")
	pool.Put("Nursuandy")
	pool.Put("Ritonga")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Done")
}
