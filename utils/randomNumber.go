package utils

import (
	"math/rand"
)

func GenerateRandom4DigitNumber() int {
	return rand.Intn(9000) + 1000
}
