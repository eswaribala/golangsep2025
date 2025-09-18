package main

import (
	"log"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/eswaribala/claimapp/healthinsurance/interfaces"
	"github.com/eswaribala/claimapp/healthinsurance/store"
)

// common catch function to recover from panic
func RecoverFromPanic() {
	if r := recover(); r != nil {
		log.Println("Recovered in RecoverFromPanic:", r)
	}
}

func main() {
	defer RecoverFromPanic()
	var memberRepo interfaces.MemberRepo = nil
	memberInstance := store.Member{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Phone:     gofakeit.Phone(),
		Address:   gofakeit.Address().Address,
		StartDate: gofakeit.Date().Format("2006-01-02"),
		EndDate:   gofakeit.Date().Format("2006-01-02"),
	}
	memberRepo = &memberInstance
	success, err := memberRepo.SaveMember()
	if err != nil {
		panic(err)
	}
	if success {
		println("Member saved successfully")
	}

	// Fetch all members
	members, err := memberRepo.GetAllMembers()
	if err != nil {
		panic(err)
	}
	for _, m := range members {
		println("Member:", m.FirstName, m.LastName, m.Email)
	}

	// Fetch a member by ID
	member, err := memberRepo.GetMemberByID(2)
	if err != nil {
		panic(err)
	}
	println("Fetched Member by ID:", member.FirstName, member.LastName, member.Email)
	// Update member's email and phone
	updated, err := memberRepo.UpdateMember("param@example.com", "123-456-7890")
	if err != nil {
		panic(err)
	}
	if updated {
		println("Member updated successfully")
	}
	// Delete a member by ID
	/*
		deleted, err := memberRepo.DeleteMember(1)

		if err != nil {
			panic(err)
		}
		if deleted {
			println("Member deleted successfully")
		}
	*/

	//find by email
	memberByEmail, err := memberRepo.GetMemberByEmail("newemail@example.com")
	if err != nil {
		panic(err)
	}
	println("Fetched Member by Email:", memberByEmail.FirstName, memberByEmail.LastName, memberByEmail.Email)

	//find by phone
	memberByPhone, err := memberRepo.GetMemberByPhone("123-456-7890")
	if err != nil {
		panic(err)
	}
	for _, m := range *memberByPhone {
		println("Fetched Member by Phone:", m.FirstName, m.LastName, m.Email)
	}
}
