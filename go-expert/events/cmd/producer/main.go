package main

import "github.com/d4niells/goutils/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "I want a large pizza, please!", "amq.direct")
}
