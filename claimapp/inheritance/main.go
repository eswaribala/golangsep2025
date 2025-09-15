package main

import (
	"fmt"
	"github.com/eswaribala/claimapp/models"
)

func main() {
	// Application entry point

	fmt.Println("Inheritance example in Go")
	// Here you can create instances of the structs defined in models/user.go
	// and demonstrate their usage.
	Admin := models.Admin{
		User:       models.User{Username: "adminUser", Password: "securePass"},
		AssignRole: true,
	}
	fmt.Printf("Admin: %+v\n", Admin)

	ClaimOfficer := models.Claimofficer{
		User:          models.User{Username: "claimOfficer", Password: "officerPass"},
		ClaimsHandled: 5,
		ProcessClaims: true,
	}
	fmt.Printf("Claim Officer: %+v\n", ClaimOfficer)
}
