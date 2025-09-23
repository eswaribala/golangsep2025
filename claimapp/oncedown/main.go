package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func StopAll() {

	fmt.Println("Processing stop all...")
	time.Sleep(2 * time.Second)
	fmt.Println("Stop all processed.")
}

func PremiumPayment(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Processing premium payment...")
	time.Sleep(2 * time.Second)
	fmt.Println("Premium payment processed.")

}

func ReportClaim(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Processing report claim...")
	time.Sleep(2 * time.Second)
	fmt.Println("Report claim processed.")
	once.Do(StopAll)

}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go PremiumPayment(&wg)
	go ReportClaim(&wg)
	wg.Wait()
	fmt.Println("All tasks completed.")

}
