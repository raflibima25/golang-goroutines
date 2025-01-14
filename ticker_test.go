package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	// using goroutine
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for resultTicker := range ticker.C {
		fmt.Println(resultTicker)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for resultTicker := range channel {
		fmt.Println(resultTicker)
	}
}
