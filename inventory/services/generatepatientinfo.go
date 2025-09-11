package services

import (
	"inventory/models"
	"math/rand"

	"github.com/bxcodec/faker/v4"
)

// pointer receiver
func GeneratePatientInfo() *models.Patient {

	// Generate and return patient information
	//Structure access using pointer
	var patient = &models.Patient{}

	patient.FirstName = faker.FirstName()
	patient.LastName = faker.LastName()
	patient.Email = faker.Email()
	patient.Phone = faker.Phonenumber()
	patient.DOB = &models.Date{
		Day:   rand.Intn(28) + 1, // to avoid invalid dates
		Month: rand.Intn(11) + 1,
		Year:  rand.Intn(50) + 1950,
	}

	return patient
}
