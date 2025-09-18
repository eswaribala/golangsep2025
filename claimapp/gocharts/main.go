package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/eswaribala/claimapp/gocharts/store"
)

func main() {

	models := []string{"BMW", "AUDI", "HYUNDAI", "KIA", "FORD"}
	claims := []store.Claim{}

	//generate 50 claims with random amounts
	for i := 0; i < 25; i++ {
		claim := store.Claim{
			ID:        uint(i + 1),
			ModelName: models[gofakeit.Number(0, len(models)-1)], // Random model
			Amount:    gofakeit.IntRange(100000000, 1000000000),  // Random amount
		}
		claims = append(claims, claim)
	}

	store.GenerateBarGraphClaims(claims, models)
}
