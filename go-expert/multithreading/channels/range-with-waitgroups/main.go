package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publisher(ch)
	go consumer(ch, &wg)

	wg.Wait()
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for msg := range ch {
		fmt.Printf("Reading message: %d\n", msg)
		wg.Done()
	}
}

func publisher(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
