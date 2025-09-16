package main

import "github.com/brianvoe/gofakeit/v7"

func validate(min int, max int, otp func(int, int) int) int {
	if min <= max && min != 0 && max != 0 {
		return otp(min, max)
	} else {
		return 0
	}
}

func main() {
	//passing anonymous function as parameter
	// Anonymous function to generate a random OTP between min and max
	otp := validate(1000, 9999, func(min int, max int) int {
		return gofakeit.IntRange(min, max)
	})
	println(otp)
}
