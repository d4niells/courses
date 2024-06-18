package main

import (
	"fmt"
	"time"
)

const (
	NUM_WORKERS = 100
	NUM_DATA    = 1000
)

func worker(workerID int, data <-chan int) {
	for x := range data {
		fmt.Printf("worker %d received %d\n", workerID, x)
		time.Sleep(1 * time.Second)
	}
}

func sendData(data chan<- int) {
	for i := 0; i < NUM_DATA; i++ {
		data <- i
	}
}

func main() {
	data := make(chan int)

	for i := 0; i < NUM_WORKERS; i++ {
		go worker(i, data)
	}

	sendData(data)
}
