package learn_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"golang.org/x/exp/rand"
)

// Struct untuk menyimpan data transaksi
type Transaction struct {
	ID     int
	Name   string
	Amount float64
}

// Fungsi untuk memproses pembayaran (simulasi)
func processPayment(workerID int, transaction Transaction, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Tandai worker selesai saat keluar

	// Simulasi waktu proses pembayaran (1-3 detik)
	processTime := time.Duration(rand.Intn(3)+1) * time.Second
	fmt.Printf("💳 Worker %d sedang memproses transaksi #%d untuk %s (Rp%.2f)...\n",
		workerID, transaction.ID, transaction.Name, transaction.Amount)
	time.Sleep(processTime)

	// Kirim hasil transaksi
	results <- fmt.Sprintf("✅ Worker %d selesai memproses transaksi #%d (Rp%.2f) untuk %s!",
		workerID, transaction.ID, transaction.Amount, transaction.Name)
}

func TestPayment(t *testing.T) {
	rand.New(rand.NewSource(uint64(time.Now().UnixNano()))) // Seed agar random berbeda tiap kali jalan

	// Daftar transaksi
	transactions := []Transaction{
		{1, "Adiva", 50000},
		{2, "Budi", 75000},
		{3, "Citra", 120000},
		{4, "Doni", 30000},
		{5, "Eka", 99000},
		{6, "Fajar", 150000},
		{7, "Gita", 25000},
		{8, "Hadi", 70000},
		{9, "Irfan", 80000},
		{10, "Joko", 65000},
		{11, "Kiki", 110000},
		{12, "Lina", 45000},
		{13, "Mira", 130000},
		{14, "Nando", 90000},
		{15, "Olivia", 55000},
	}

	totalWorkers := 5                      // Jumlah worker yang berjalan paralel
	jobs := make(chan Transaction, len(transactions)) // Channel untuk daftar transaksi
	results := make(chan string, len(transactions))   // Channel untuk hasil pembayaran
	var wg sync.WaitGroup

	// Membuat worker pool
	for i := 1; i <= totalWorkers; i++ {
		go func(workerID int) {
			for transaction := range jobs {
				processPayment(workerID, transaction, results, &wg)
			}
		}(i)
	}

	// Mengirimkan transaksi ke channel jobs
	for _, transaction := range transactions {
		wg.Add(1)
		jobs <- transaction
	}
	close(jobs) // Setelah semua transaksi masuk, tutup channel jobs

	// Menunggu semua transaksi selesai
	wg.Wait()
	close(results) // Tutup channel hasil

	// Menampilkan hasil transaksi
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("🎉 Semua transaksi telah diproses!")
}