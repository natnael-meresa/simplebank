package utils

import (
	"math/rand"
	"strings"
	"time"
)

// go hack
// think like programmer
// kira go book
// cockroach
// microservice

// bdd book
const alphabets = "abcdefghijklmnopkrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var randString strings.Builder
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(len(alphabets))]

		randString.WriteByte(c)
	}

	return randString.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomAmount() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}

	return currencies[rand.Intn(len(currencies))]
}

func RandomID() int64 {
	return RandomInt(1, 100)
}
