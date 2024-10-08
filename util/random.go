package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))

}

// RandomInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generate an owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMony generate a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 100000)
}

// RandomCurrency generate a random currency code
func RandomCurrency() string {
	currencies := []string{"SAR", "AUE", "QAR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
