package main

import (
	"math/rand"

	"github.com/eswaribala/claimapp/models"
)

func main() {

	//create claim amount
	/*
		rootClaim := &models.Claim{
			ID:     1,
			Amount: 100,
			Claims: []*models.Claim{
				{
					ID:     2,
					Amount: 200,
				},
				{
					ID:     3,
					Amount: 300,
				},
				{
					ID:     4,
					Amount: 400,
				},
			},
		}
	*/

	rootClaim := &models.Claim{
		ID:     1,
		Amount: rand.Intn(400) + 100,
	}

	//generate random claims
	for i := 0; i < 5; i++ {
		claim := &models.Claim{
			ID:     uint(i + 2),
			Amount: rand.Intn(400) + 100,
		}
		rootClaim.Claims = append(rootClaim.Claims, claim)
	}

	total := rootClaim.TotalClaimsAmount()
	println("Total Claims Amount:", total) // Output: Total Claims Amount: 1000
}
