package learn_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			// OnlyOnce()
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
	fmt.Println("Done")
}
