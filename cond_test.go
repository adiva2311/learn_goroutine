package learn_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

// Untuk mengirim signal ke cond.Wait() untuk menjalankan perintah satu per satu
func SignalCond() {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		cond.Signal()
	}
}

// Untuk mengirim signal ke cond.Wait() untuk menjalankan perintah langsung semua
func BroadCond() {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}
}

func TestCond(t *testing.T) {
	for i := 1; i <= 10; i++ {
		go WaitCondition(i)
	}

	//go SignalCond()
	go BroadCond()

	group.Wait()
}
