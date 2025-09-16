package main

import (
	"fmt"

	"io"
	"net/http"

	"github.com/brianvoe/gofakeit/v7"
)

// common catch function to recover from panic
func RecoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in RecoverFromPanic:", r)
	}
}

func main() {

	otp := validate(10000, 9999, func(min int, max int) int {
		return gofakeit.IntRange(min, max)
	})
	println("OTP is:", otp)
	age := ComputeAge(120) //panic occurs here
	fmt.Println("Age is:", age)
	response := GetResponse("https://jsonplaceholder.typicode.com/users")
	fmt.Println("Response is:", response)

	fmt.Println("Program continues...")

}

func ComputeAge(days int) int {
	defer RecoverFromPanic()
	currentYear := 2025
	fixedDays := 0
	birthYear := (days / fixedDays) + 2000
	println(birthYear)
	return currentYear - birthYear
}

func validate(min int, max int, otp func(int, int) int) int {
	defer RecoverFromPanic()
	if min <= max && min != 0 && max != 0 {
		return otp(min, max)
	} else {
		panic("Invalid min or max values")
	}
}

func GetResponse(url string) string {
	defer RecoverFromPanic()
	// Simulate a network call
	if url == "" {
		panic("Invalid URL")
	}
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("HTTP request failed with status: %s", response.Status))
	}
	// Simulate processing the response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return "Response from " + url + ": " + string(body)
}
