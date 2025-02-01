package learn_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Untuk mengatasi Race Condition di dalam GoRoutine kita menggunakan sync.Mutex
// Dan sebelum dan sesudah variabel yang ingin dirubah harus kita lock dan unlock
// Kapan menggunakan mutex adalah apabila ada sebuah variabel yang digunakan secara sharing/banyak goroutine yang menjalankannya

func TestMutex(t *testing.T) {
	counter := 0

	var mutex sync.Mutex

	for x := 1; x <= 1000; x++ {
		go func() {
			for y := 1; y <= 100; y++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", counter)
}

// sync.RWMutex -> Read n Write Mutex
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	// Untuk menulis atau Merubah (WRITE)
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	// Untuk mengambil atau Melihat (READ)
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()

	return balance
}

func TestReadWriteMutes(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance =", account.GetBalance())
}

// Test Deadlock
type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, group *sync.WaitGroup, amount int) {
	defer group.Done()

	// **Hindari deadlock dengan menentukan urutan penguncian berdasarkan alamat memori**
	var first, second *UserBalance
	if user1.Name < user2.Name { // Urutkan berdasarkan nama untuk konsistensi
		first, second = user1, user2
	} else {
		first, second = user2, user1
	}

	first.Lock()
	fmt.Println("Lock User 1", first.Name)
	first.Change(-amount)
	defer first.Unlock()

	second.Lock()
	fmt.Println("Lock User 2", second.Name)
	second.Change(amount)
	defer second.Unlock()

}

func TestDeadlock(t *testing.T) {
	var group sync.WaitGroup

	user1 := UserBalance{
		Name:    "Adiva",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Ritonga",
		Balance: 1000000,
	}

	group.Add(2)
	go Transfer(&user1, &user2, &group, 200000)
	go Transfer(&user2, &user1, &group, 100000)

	group.Wait()
	//time.Sleep(3 * time.Second)
	fmt.Println("User", user1.Name, "Balance", user1.Balance)
	fmt.Println("User", user2.Name, "Balance", user2.Balance)
}
