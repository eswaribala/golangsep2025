package main

import (
	"fmt"

	"github.com/eswaribala/claimapp/models"
)

func main() {
	//create policy instance
	policy := models.Policy{
		Name: "Health Insurance",
		Coverage: struct {
			MaxAmount float64 `gorm:"not null" json:"max_amount"`
			MinAmount float64 `gorm:"not null" json:"min_amount"`
		}{
			MaxAmount: 100000.0,
			MinAmount: 5000.0,
		},
	}
	fmt.Printf("Policy: %+v\n", policy)
}
