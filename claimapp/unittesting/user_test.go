package main

import (
	"encoding/csv"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// open csv file and read data
func OpenCSVFile(filePath string) (*csv.Reader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	//defer file.Close()

	reader := csv.NewReader(file)
	return reader, nil
}

func TestValidateUser(t *testing.T) {
	userInstance := User{
		Name:     "JohnDoe",
		Email:    "john.doe@example.com",
		Password: "P@ssw0rd",
	}

	assert.NoError(t, ValidateUser(userInstance))
}

func TestUserFromCSV(t *testing.T) {
	csvReader, err := OpenCSVFile("userdata.csv")
	if err != nil {
		t.Fatalf("Failed to open CSV file: %v", err)
	}
	records, err := csvReader.ReadAll()
	if err != nil {
		t.Fatalf("Failed to read CSV file: %v", err)
	}
	for _, record := range records[1:] {
		user := User{
			Name:     record[0],
			Email:    record[1],
			Password: record[2],
		}
		log.Println("Testing user:", user)
		err := ValidateUser(user)
		assert.NoError(t, err, "Validation failed for user: %+v", user)
	}
}
