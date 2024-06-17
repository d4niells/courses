package main

import (
	"fmt"
	"sync"
	"time"
)

func Task(name string, waitGroup *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Task %s is running\n", name)
		time.Sleep(1 * time.Second)
		waitGroup.Done()
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)

	// Thread 2
	go Task("A", &waitGroup)
	// Thread 3
	go Task("B", &waitGroup)
	// Thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "Anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
}
