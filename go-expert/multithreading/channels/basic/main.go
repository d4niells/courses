package main

import "fmt"

func main() {
	msg := make(chan string)

	go func() {
		// Send a value into a channel using the sintax <-
		msg <- "Hello world"

		// It won't send the message to the channel because it's not empty.
		// The channel must be read before sending another value to it.
		msg <- "Hello world two"
	}()
	// Reading channel values
	fmt.Println(<-msg)
}
