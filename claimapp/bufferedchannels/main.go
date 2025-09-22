package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/eswaribala/claimapp/models"
)

func main() {

	claimChannel := make(chan models.Claim, 5)
	noOfClaims := gofakeit.IntRange(5, 15)
	//claim array
	var claims []models.Claim
	//simulate adding claims to the channel
	for i := 1; i <= noOfClaims; i++ {
		claim := models.Claim{
			ID:     uint(gofakeit.IntRange(10000, 10000000)),
			Amount: gofakeit.IntRange(10000, 10000000),
		}
		claims = append(claims, claim)

	}

	go func() {
		//claim raised by the client
		println("Raising claims...")
		count := 0
		for _, claim := range claims {
			claimChannel <- claim
			count++
			println("Claims raised:", count)
		}

	}()

	//claim processed by the server
	println("Processing claims...")
	for i := 1; i <= noOfClaims; i++ {
		claim := <-claimChannel
		println("Claim ID:", claim.ID, "Amount:", claim.Amount)
	}

}
