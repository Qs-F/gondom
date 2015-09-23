package gondom

import (
	"math/rand"
)

const (
	_rawURLLetters = "abcdefghijklmnopqrstuvwxyz0123456789"
	_rawLetters    = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Make(n int, seed int64) string {
	b := make([]byte, n)
	rand.Seed(seed)
	for i := range b {
		b[i] = _rawLetters[rand.Intn(62)]
	}
	return string(b)
}

func MakeURL(n int, seed int64) string {
	b := make([]byte, n)
	rand.Seed(seed)
	for i := range b {
		b[i] = _rawURLLetters[rand.Intn(36)]
	}
	return string(b)
}
