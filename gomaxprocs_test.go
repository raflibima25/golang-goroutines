package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("total CPU:", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1) // gunakan -1 agar tidak mengubah jika menggunakan lebih dari 0
	fmt.Println("total thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine:", totalGoroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("total CPU:", totalCPU)

	runtime.GOMAXPROCS(20)                // untuk ubah Thread Golang
	totalThread := runtime.GOMAXPROCS(-1) // gunakan -1 agar tidak mengubah jika menggunakan lebih dari 0
	fmt.Println("total thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine:", totalGoroutine)

	group.Wait()
}
