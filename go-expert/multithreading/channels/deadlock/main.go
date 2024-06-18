package main

func main() {
	forever := make(chan bool)

	// The go runtime was waiting for the channel to be filled to empty it,
	// but that's never happen, and then it will result in a deadlock error.
	<-forever
}
