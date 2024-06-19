package main

import (
	"fmt"
	"time"
)

type Message struct {
	ID    int
	Value string
}

func kafkaPub(ch chan Message) {
	for {
		time.Sleep(time.Second)
		msg := Message{1, "Hello from kafka"}
		ch <- msg
	}
}

func rabbitMQPub(ch chan Message) {
	for {
		time.Sleep(2 * time.Second)
		msg := Message{2, "Hello from rabbitMQ"}
		ch <- msg
	}
}

func main() {
	ch1 := make(chan Message)
	ch2 := make(chan Message)

	go kafkaPub(ch1)
	go kafkaPub(ch2)

	for {
		select {
		case msg := <-ch1:
			fmt.Printf("Received from Kafka: %s\n", msg.Value)
		case msg2 := <-ch2:
			fmt.Printf("Received from RabbitMQ: %s\n", msg2.Value)
		case <-time.After(3 * time.Second):
			println("There're no messages to consume")
			// default: // We got default if all channels are empty
			// 	println("default")
		}
	}
}
