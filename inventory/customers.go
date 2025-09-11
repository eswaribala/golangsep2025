package main

import (
	"github.com/bxcodec/faker/v4"
)

func GenerateCustomers() []string {
	//slice of customers
	//make function creates a slice of given length and capacity
	customers := make([]string, 50)
	for i := range customers {
		customers[i] = faker.FirstName()
	}
	//adding more customers to the slice
	customers = append(customers, faker.FirstName())
	customers = append(customers, faker.FirstName())
	customers = append(customers, faker.FirstName())
	return customers
}
