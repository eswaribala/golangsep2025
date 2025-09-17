package main

import (
	"github.com/eswaribala/claimapp/healthinsurance/interfaces"
	"github.com/eswaribala/claimapp/healthinsurance/store"
)

func main() {
	var memberRepo interfaces.MemberRepo = nil
	memberInstance := store.Member{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     "123-456-7890",
		Address:   "123 Main St, Anytown, USA",
		StartDate: "2020-01-01",
		EndDate:   "2020-12-31",
	}
	memberRepo = &memberInstance
	success, err := memberRepo.SaveMember()
	if err != nil {
		panic(err)
	}
	if success {
		println("Member saved successfully")
	}

}
