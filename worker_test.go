package learn_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Fungsi yang akan dijalankan oleh worker
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// Simulasi kerja (misalnya proses butuh waktu 1-3 detik)
		time.Sleep(time.Duration(1+job%3) * time.Second)
		results <- fmt.Sprintf("Worker %d selesai mengerjakan tugas %d", id, job)
	}
}

func TestCook(t *testing.T) {
	totalJobs := 10   // Jumlah tugas yang harus dikerjakan
	totalWorkers := 3 // Jumlah worker yang tersedia

	jobs := make(chan int, totalJobs)    // Channel untuk menampung tugas
	results := make(chan string, totalJobs) // Channel untuk menerima hasil
	var wg sync.WaitGroup

	// Membuat Worker Pool (Menjalankan Worker)
	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Mengirimkan tugas ke channel jobs
	for j := 1; j <= totalJobs; j++ {
		jobs <- j
	}
	close(jobs) // Menutup channel jobs agar worker tahu tidak ada tugas baru

	// Menunggu semua worker selesai
	wg.Wait()
	close(results) // Menutup hasil setelah semua tugas selesai

	// Menerima dan mencetak hasil kerja worker
	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("✅ Semua tugas telah selesai!")
}