package learn_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// Race Condition adalah dimana ada goroutine yang berjalan secara bersamaan dengan satu variabel
func TestRaceCondition(t *testing.T) {
	counter := 0

	for x := 1; x <= 1000; x++ {
		go func() {
			for y := 1; y <= 100; y++ {
				counter++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", counter)
}
