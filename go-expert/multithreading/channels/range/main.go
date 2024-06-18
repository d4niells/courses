package main

import "fmt"

func main() {
	ch := make(chan int)
	go publisher(ch)
	reader(ch)
}

func reader(ch chan int) {
	for msg := range ch {
		fmt.Printf("Reading message: %d\n", msg)
	}
}

func publisher(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
