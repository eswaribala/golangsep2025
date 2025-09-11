package models

type PatientInfo interface {
	ShowPatientInfo(newMail string)
	DisplayPatientInfo(newMail string)
}

// call by value
func (patient Patient) ShowPatientInfo(newMail string) {
	println("Patient Information:")
	println("First Name:", patient.FirstName)
	println("Last Name:", patient.LastName)
	println("Date of Birth:", patient.DOB.Day, "/", patient.DOB.Month, "/", patient.DOB.Year)
	patient.Email = newMail // Modifying the email field
	println("Email:", patient.Email)
	println("Phone:", patient.Phone)
}

//call by reference

func (patient *Patient) DisplayPatientInfo(newMail string) {
	println("Patient Information:")
	println("First Name:", patient.FirstName)
	println("Last Name:", patient.LastName)
	println("Date of Birth:", patient.DOB.Day, "/", patient.DOB.Month, "/", patient.DOB.Year)
	patient.Email = newMail // Modifying the email field
	println("Email:", patient.Email)
	println("Phone:", patient.Phone)
}
