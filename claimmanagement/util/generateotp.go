package util

import "math/rand"

func GenerateOTP() int64 {

	return rand.Int63n(899999) + 100000

}
