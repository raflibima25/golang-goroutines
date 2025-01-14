package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter", x)
}

/*
	Masalah Race Condition
	jadi ada beberapa goroutine yang counter (menambah) dengan value yang sama-
	karena sifatnya yang concurency dan parallel programming
	contoh 5 goroutine:
		x = 1000 + 1
		x = 1000 + 1
		x = 1000 + 1
		x = 1000 + 1
		x = 1000 + 1
	goroutine ini sama2 menambah data dengan value yang sama
*/
