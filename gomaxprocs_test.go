package learn_goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T) {
	totalCPU := runtime.NumCPU()
	fmt.Println(totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println(totalGoroutine)
}
