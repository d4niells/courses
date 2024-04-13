package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(3)*time.Second)
	defer cancel()

	Do(ctx)
}

func Do(ctx context.Context) {
	select {
	case <-ctx.Done():
		println("Done!")
	case <-time.After(2 * time.Second):
		println("Canceled")
	}
}
