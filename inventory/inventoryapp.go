package main

import (
	"fmt"
	"inventory/models"
	"inventory/services"
	"inventory/utility"
	"math/rand"

	"github.com/bxcodec/faker/v4"
)

// global variable (package level variable)
var organization = "NIIT"
var host string
var port int
var user string
var passwordValue string

func init() {
	host = "localhost"
	port = 5432
	user = faker.FirstName()
	passwordValue = faker.Password()
	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
	fmt.Println("User:", user)
	fmt.Println("Password:", passwordValue)
}

func main() {

	host = "172.68.91.10"
	fmt.Println("Updated Host:", host)
	fmt.Println(utility.GenerateDeviceInfo())
	//fmt.Println("Location:", location)
	invoke()
	//ReadData()
	for _, donor := range GenerateDonors() {
		fmt.Println("Donor:", donor)
	}

	for index, customer := range GenerateCustomers() {
		fmt.Println("Customer", index+1, ":", customer)
	}

	/*
		//call generate plans function from services package
		for id, plan := range services.GeneratePlans() {
			fmt.Printf("Plan ID: %d, Plan Name: %s\n", id, plan)
		}
	*/
	totalPayment := services.AggregatePayments(100, 200, 300, 400, 500)
	fmt.Println("Total Payment:", totalPayment)
	totalPayment += services.AggregatePayments(100, 200, 300, 400, 500, 4387, 48070, 40700)
	fmt.Println("Total Payment:", totalPayment)

	test()

	//destructure using map
	//anonymous structure for json or csv
	patient := services.GeneratePatientInfo()

	patientMap := services.StructToMapPatient(patient)
	fmt.Println("Patient Structure to Map using Mapstructure:")
	for key, value := range patientMap {
		if key == "DOB" {
			dateMap := services.StructToMapDate(patient.DOB)
			fmt.Println("DOB Details:")
			for dateKey, dateValue := range dateMap {
				fmt.Printf("%s : %v\n", dateKey, dateValue)
			}
		} else {
			fmt.Printf("%s : %v\n", key, value)
		}
	}

	/*
		//reading structure using pointer
		fmt.Printf("Patient Email: %s\n	", patient.Email)
		fmt.Printf("Patient Name: %s %s\n", patient.FirstName, patient.LastName)
		fmt.Printf("Patient DOB: %02d-%02d-%d\n", patient.DOB.Day, patient.DOB.Month, patient.DOB.Year)
		fmt.Printf("Patient Phone: %s\n", patient.Phone)
	*/

	//instance of the structure
	patientv1 := &models.Patient{}
	patientv1.FirstName = faker.FirstName()
	patientv1.LastName = faker.LastName()
	patientv1.Email = faker.Email()
	patientv1.Phone = faker.Phonenumber()
	patientv1.DOB = &models.Date{
		Day:   rand.Intn(28) + 1, // to avoid invalid dates
		Month: rand.Intn(11) + 1,
		Year:  rand.Intn(50) + 1950,
	}
	//call by value
	patientv1.ShowPatientInfo("param@example.com")
	//call by reference
	patientv1.DisplayPatientInfo("eswari@example.com")

	//interface
	var patientInfo models.PatientInfo = nil
	patientInfo = patientv1
	patientInfo.ShowPatientInfo("interface@example.com")
	patientInfo.DisplayPatientInfo("interface@example.com")

}

func invoke() {
	// local variable (function level variable)
	var location = "Bangalore"
	fmt.Println("Organization:", organization)
	fmt.Println("Location:", location)

}

func test() {
	var amount int = 10000
	fmt.Println("Amount:", amount)
	//pointer
	var ptr *int = &amount
	fmt.Println("Pointer address:", ptr)
	fmt.Println("Pointer value:", *ptr)
	amount = rand.Intn(1000000)
	fmt.Println("Updated Amount:", amount)
	fmt.Println("Pointer value after amount update:", *ptr)
}
