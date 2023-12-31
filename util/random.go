package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	numbers  = "0123456789"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail(username string) string {
	return username + "@email.com"
}

func RandomType() string {
	n := RandomInt(0, 1)

	if n == 0 {
		return "user"
	}

	return "store"
}

func RandomDocument(userType string) string {
	n := 14

	if userType == "user" {
		n = 11
	}

	var sb strings.Builder
	k := len(numbers)

	for i := 0; i < n; i++ {
		c := numbers[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
