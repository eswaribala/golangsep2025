package main

import "github.com/brianvoe/gofakeit/v7"

//closures are functions that reference variables from outside their body

func claimClosure() func() int {
	basicClaimAmount := 10000
	return func() int {
		basicClaimAmount += int(gofakeit.Float32Range(0.10, 0.50) * float32(basicClaimAmount))
		return basicClaimAmount
	}
}

func main() {
	claim := claimClosure()
	println(claim())
	println(claim())
	println(claim())
}
