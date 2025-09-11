package main

import (
	"github.com/bxcodec/faker/v4"
)

func GenerateDonors() []string {

	var donors [10]string
	for i := 0; i < 10; i++ {

		donors[i] = faker.Name()
	}
	return donors[:]
}
