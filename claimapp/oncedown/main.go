package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var dbConnection string

func InitDB() {

	fmt.Println("Database Initialized")
	dbConnection = "DB Connection Established"

}

func getDBConnection() string {
	once.Do(InitDB)
	return dbConnection
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			conn := getDBConnection()
			fmt.Printf("Goroutine %d: %s\n", id, conn)
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines finished.")
}
