package helpers

import (
	"math/rand"
	"time"
)

func GenerateRandomString() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < 6; i++ {
		randomIndex := r.Intn(len(charSet))
		result += string(charSet[randomIndex])
	}
	return result
}
