package main

import (
	"fmt"
	"sync"
	"time"
)

func Enrollment(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Enrollment started")
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Println("Enrollment completed")
}

func DeEnrollment(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("DeEnrollment started")
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Println("DeEnrollment completed")
}

func Billing(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Billing started")
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Println("Billing completed")
}

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	go Enrollment(&wg)

	go Billing(&wg)

	go DeEnrollment(&wg)

	wg.Wait()

	fmt.Println("All tasks completed")
}
