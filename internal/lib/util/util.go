package util

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")

	rndStr := make([]rune, size)
	for i := range rndStr {
		rndStr[i] = chars[rnd.Intn(len(chars))]
	}

	return string(rndStr)
}
