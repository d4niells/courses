package main

import (
	"fmt"
	"time"
)

func Task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Task %s is running\n", name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	// Thread 2
	go Task("A")
	// Thread 3
	go Task("B")
	// Thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "Anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second)
}
