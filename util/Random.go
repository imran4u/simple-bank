package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var r *rand.Rand

func init() {
	// rand.Seed(time.Now().UnixNano()) // Deprecated in Go 1.20+

	// Create a new random number generator with a seed
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

}

func RandomInt(min, max int64) int64 {
	return min + r.Int63n(max-min+1)
}

func RandomString(n int) string {
	var builder strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		b := alphabet[r.Intn(k)]
		builder.WriteByte(b)
	}
	return builder.String()
}

func RandomName() string {
	return RandomString(6)
}

func RandomAmount() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{
		USD,
		RS,
	}
	return currencies[r.Intn(len(currencies))]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
