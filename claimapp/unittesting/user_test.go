package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	userInstance := User{
		Name:     "JohnDoe",
		Email:    "john.doe@example.com",
		Password: "P@ssw0rd",
	}

	assert.NoError(t, ValidateUser(userInstance))
}
