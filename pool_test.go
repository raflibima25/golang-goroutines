package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any { // jika ingin set default value Pool, karena jika data pool kosong akan return nil
			return "New"
		},
	}

	// menyimpan data ke pool
	pool.Put("Rafli")
	pool.Put("Bima")
	pool.Put("Pratandra")

	for i := 0; i < 10; i++ {
		go func() {
			// .Get() mengambil data di pool
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")

}
