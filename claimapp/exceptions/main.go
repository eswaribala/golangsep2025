package main

import (
	"errors"

	"github.com/brianvoe/gofakeit/v7"
)

func claimClosure() func() (int, error) {
	basicClaimAmount := gofakeit.IntRange(0, 10000)
	return func() (int, error) {
		basicClaimAmount += int(gofakeit.Float32Range(0.10, 0.50) * float32(basicClaimAmount))
		if basicClaimAmount == 0 {
			return 0, errors.New("claim amount is zero")
		} else {
			return basicClaimAmount, nil
		}
	}
}

func main() {
	claim := claimClosure()
	amount, err := claim()
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Claim Amount:", amount)
	}
}
