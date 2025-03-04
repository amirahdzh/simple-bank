package util

import (
	"strings"
	"time"

	"math/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // min->max
}

// RandomString generates a random string with length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 2000)
}

// RandomCurrency generates a randon currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "SAR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
