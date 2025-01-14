package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Mutex
// Ini adalah solusi meng-handle dari masalah Race Condition di Golang
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter", x)
}

// sync.RWMutex
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock() // write lock
	account.Balance += amount
	account.RWMutex.Unlock() // write unlock
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock() // read lock
	balance := account.Balance
	account.RWMutex.RUnlock() // read unlock
	return balance
}

func TestRWMutex(t *testing.T) {
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
	fmt.Println("Total Balance:", account.Balance)
}

// Deadlock
type UserBalance struct {
	sync.Mutex
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
	user.Balance += amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	fmt.Println("Unlock user1", user1.Name)
	user2.Unlock()
	fmt.Println("Unlock user2", user2.Name)
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Bima",
		Balance: 10000,
	}

	user2 := UserBalance{
		Name:    "Raka",
		Balance: 10000,
	}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 2000)

	time.Sleep(10 * time.Second)

	fmt.Println("User1", user1.Name, "balance:", user1.Balance)
	fmt.Println("User2", user2.Name, "balance:", user2.Balance)
}
