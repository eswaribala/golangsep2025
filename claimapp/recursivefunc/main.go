package main

import "github.com/eswaribala/claimapp/models"

func main() {

	//create claim amount
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

	total := rootClaim.TotalClaimsAmount()
	println("Total Claims Amount:", total) // Output: Total Claims Amount: 1000
}
