package learn_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// Gunakan RNG global agar lebih aman
//var rng = rand.New(rand.NewSource(uint64(time.Now().UnixNano())))

// Kapasitas tiket yang tersedia
var totalTickets = 200
var ticketMutex sync.Mutex // Mutex untuk menghindari race condition

// Struct untuk menyimpan data pembelian tiket
type TicketOrder struct {
	ID   int
	Name string
}

// Fungsi untuk memproses pembelian tiket
func processTicket(workerID int, order TicketOrder, results chan<- string, wg *sync.WaitGroup){
	defer wg.Done()

	// Simulasi waktu proses pembayaran (1-3 detik)
	// processTime := time.Duration(rand.Intn(3)+1) * time.Second
	// time.Sleep(processTime)

	//ticketMutex.Lock()
	if totalTickets > 0 {
		totalTickets--
		results <- fmt.Sprintf("✅ Worker %d : %s BERHASIL membeli tiket. Sisa Tiket : %d", workerID, order.Name, totalTickets)
	
	} else {
		results <- fmt.Sprintf("❌ Worker %d : %s GAGAL membeli tiket. Tiket Sudah Habis!!", workerID, order.Name)
	}
	//ticketMutex.Unlock()
}

func TestTicket(t *testing.T) {
	// Memasukkan 300 orang yang membeli tiket
	orders := []TicketOrder{}
	for i := 1; i <= 300; i++{
		orders = append(orders, TicketOrder{
			ID: i,
			Name: fmt.Sprintf("Pembeli ke %d", i),
		})
	}

	totalWorkers := 5
	jobs := make(chan TicketOrder, len(orders)) // Channel untuk daftar pembeli tiket
	results := make(chan string, len(orders))	// Channel untuk hasil transaksi

	var wg sync.WaitGroup

	// Function Workers
	for i := 1; i <= totalWorkers; i++ {
		go func(workerID int) {
			for order := range jobs {
				processTicket(workerID, order, results, &wg)
			}
		}(i)
	}

	// Mengirim order ke channel
	for _, order := range orders{
		wg.Add(1)
		jobs <- order
	}
	close(jobs)

	wg.Wait()
	close(results)

	// Menampilkan semua hasil result
	for result := range results{
		fmt.Println(result)
	}

	fmt.Println("SELESAI!!!!")
}