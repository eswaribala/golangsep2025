package utility

import (
	"github.com/bxcodec/faker/v4"
)

// GenerateDeviceInfo generates a random MAC address for a device.
// External functions, variables start with capital letters
func GenerateDeviceInfo() string {
	return faker.MacAddress() + "-" + generateID()
}
