package main

import "time"

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()

	select {
	case msg1 := <-ch1:
		println("received:", msg1)
	case msg2 := <-ch2:
		println("received:", msg2)
	case <-time.After(3 * time.Second):
		println("timeout")
		// default: // We got default if all channels are empty
		// 	println("default")
	}
}
