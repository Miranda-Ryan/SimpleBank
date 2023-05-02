package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomOwner generates a random string of length n as owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomCurrency generates a random currency from the list of currencies
func RandomCurrency() string {
	currencies := []string{"USD", "CAD", "EUR"}
	currency := currencies[rand.Intn(len(currencies))]

	return currency
}

// RandomString generates a random string based on n length
func RandomString(n int) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		randomInt := rand.Intn(len(alphabet))
		randomChar := alphabet[randomInt]
		builder.WriteByte(randomChar)
	}

	return builder.String()
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}
