package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)

	go doAutomatedCancellation(ctx)

	// Simulate some work
	time.Sleep(time.Second * 10)

	fmt.Println("Main function completed")

}

func doAutomatedCancellation(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// Handle cancellation
			log.Println("Operation cancelled:", ctx.Err())
			return
		default:
			// Perform regular work
			log.Println("Working...")
			time.Sleep(time.Second * 1)
		}
	}
}
