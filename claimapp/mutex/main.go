package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eswaribala/claimapp/models"
)

var memberInstance = models.Member{
	ID:                   gofakeit.IntRange(1, 1000),
	Username:             gofakeit.Name(),
	Email:                gofakeit.Email(),
	Password:             gofakeit.Password(true, true, true, true, false, 12),
	LowIncome:            gofakeit.Bool(),
	PreferredContact:     gofakeit.RandomString([]string{"email", "phone", "text"}),
	Consent:              gofakeit.Bool(),
	EmergencyContactName: gofakeit.Name(),
}

var mutex sync.Mutex

func CallCenterApp(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Call Center App started")
	// Simulate work with sleep
	time.Sleep(2 * time.Second)
	mutex.Lock()
	memberInstance.EmergencyContactName = gofakeit.Name()
	memberInstance.LowIncome = gofakeit.Bool()
	memberInstance.PreferredContact = gofakeit.RandomString([]string{"email", "phone", "text"})
	memberInstance.Consent = gofakeit.Bool()

	fmt.Println("Processing member in Call Center App:", memberInstance)

	mutex.Unlock()

	fmt.Println("Call Center App completed")

}
func WebApp(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Web App started")
	// Simulate work with sleep
	time.Sleep(2 * time.Second)
	mutex.Lock()
	memberInstance.EmergencyContactName = gofakeit.Name()
	memberInstance.LowIncome = gofakeit.Bool()
	memberInstance.PreferredContact = gofakeit.RandomString([]string{"email", "phone", "text"})
	memberInstance.Consent = gofakeit.Bool()

	fmt.Println("Processing member in Web App:", memberInstance)

	mutex.Unlock()
	fmt.Println("Web App completed")
}

func SupplementalService(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Supplemental Service started")
	// Simulate work with sleep
	time.Sleep(2 * time.Second)
	mutex.Lock()
	memberInstance.EmergencyContactName = gofakeit.Name()
	memberInstance.LowIncome = gofakeit.Bool()
	memberInstance.PreferredContact = gofakeit.RandomString([]string{"email", "phone", "text"})
	memberInstance.Consent = gofakeit.Bool()

	fmt.Println("Processing member in Supplemental Service:", memberInstance)

	mutex.Unlock()
	fmt.Println("Supplemental Service completed")

}

func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	go CallCenterApp(&wg)
	go WebApp(&wg)
	go SupplementalService(&wg)
	wg.Wait()
	fmt.Println("All services completed")

}
