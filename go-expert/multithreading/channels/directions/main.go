package main

import "fmt"

// The chan<- syntax means it's a receive-only channel.
// You cannot send messages to this channel, only receive messages from it.
func pub(ch chan<- string) {
	ch <- "Hello"
}

// The <-chan syntax means it's a send-only channel.
// You cannot receive messages from this channel, only send messages to it.
func sub(ch <-chan string) {
	fmt.Println(<-ch)
}

func main() {
	channel := make(chan string)

	go pub(channel)
	sub(channel)
}
