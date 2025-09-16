package main

import (
	"fmt"

	"github.com/eswaribala/claimapp/models"
)

func main() {
	//create policy instance
	policy := models.Policy{
		Name: "Health Insurance",
	}
	policy.Coverage.MaxAmount = 1000000
	policy.Coverage.MinAmount = 5000

	fmt.Printf("Policy: %+v\n", policy)
}
