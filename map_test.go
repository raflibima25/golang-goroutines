package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

// ini adalah cara Map tetapi aman menggunakan goroutine
// sync.Map
func AddToMap(value int, data *sync.Map, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(i, data, group)
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
