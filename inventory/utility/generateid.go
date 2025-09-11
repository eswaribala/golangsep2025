package utility

import (
	"github.com/bxcodec/faker/v4"
)

func generateID() string {
	return faker.UUIDDigit()
}
