package learn_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// Cara mengirim DATA ke CHANNEL
	// nama_channel <- data
	//channel <- "Adiva"

	// Cara mengambil CHANNEL ke DATA
	// data := <- nama_channel
	//data := <-channel

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Adiva Nursuandy Ritonga"
		fmt.Println("Sukses Mengirim Data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

	// Channel Harus Di CLOSE
	//close(channel)
}

// Channel as Parameter

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Adiva Nursuandy Ritonga"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	fmt.Println("Sukses Mengirim Data ke Channel Sebagai Parameter")

	data := <-channel
	fmt.Println(data)

	time.Sleep(3 * time.Second)

}

// Channel Bisa digunakan hanya untuk MENGIRIM(IN) atau MENERIMA(OUT) saja

// Hanya Untuk MENGIRIM(IN)
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Adiva Nursuandy Ritonga"
}

// Hanya Untuk MENERIMA(OUT)
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

// Buffered Channel -> Untuk mengirim lebih dari 1 data ke dalam channel
// channel := make(chan string, banyaknya_buffer)
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Adiva"
		channel <- "Nursuandy"
		channel <- "Ritonga"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println(cap(channel)) // Melihat berapa banyak buffered channel
	fmt.Println(len(channel)) // Melihat berapa banyak jumlah data didalam buffer

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")

}

// Range Channel -> Menerima data yang banyak / yang tidak tentu jumlah datanya HANYA dari satu(1) Channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("perulangan ke - " + strconv.Itoa(i))
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

// Select Channel -> Mengambil data dari banyak channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// Kita harus membuat sebuah kondisi kapan for harus berhenti, disini menggunakan counter
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Channel 1 :", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Channel 2 :", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

}

func TestDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// Kita harus membuat sebuah kondisi kapan for harus berhenti, disini menggunakan counter
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Channel 1 :", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Channel 2 :", data)
			counter++
		default:
			fmt.Println("Waiting for Data")
		}
		if counter == 2 {
			break
		}
	}

}
