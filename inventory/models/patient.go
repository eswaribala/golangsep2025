package models

type Date struct {
	Day   int
	Month int
	Year  int
}

type Patient struct {
	FirstName string
	LastName  string
	DOB       *Date
	Email     string
	Phone     string
}
