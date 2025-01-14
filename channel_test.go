package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Channel wajib ada pengirim dan penerima, jika tidak ada salah satunya maka akan error
func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Rafli Bima Pratandra"
		fmt.Println("Selesai mengirim data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

// Channel sebagai parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Rafli Bima Pratandra"
	//fmt.Println("Selesai mengirim data ke Channel parameter")
}

func TestChannelParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

// Channel In dan Out
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Rafli Bima Pratandra"
	fmt.Println("Selesai mengirim data ke Channel In dan Out")
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered Channel
func TestBufferedChannel(t *testing.T) {
	/*
		Jika sebelumnya tanpa buffer harus ada pengirim dan penerima akan error karena ter-block,
		jika ditambah buffer maka tidak akan ter-block lagi karena data tersebut masuk ke buffer
		jadi jika sudah mengirim tetapi tidak yang menerima itu tidak menjadi masalah, karena sudah mengirim dan disimpan di buffer
	*/
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Rafli"
		channel <- "Bima"
		channel <- "Pratandra"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	//fmt.Println("cap:", cap(channel))
	//fmt.Println("len:", len(channel))
	fmt.Println("Selesai")
}

// Range Channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}

		close(channel) // close ini penting untuk menghentikan looping agar tidak deadlock!
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

// Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

// Default Select Channel
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data...")
		}

		if counter == 2 {
			break
		}
	}
}
