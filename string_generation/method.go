package string_generation

import (
	"math/rand"
	"strings"
)

var chars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func byteGenerateString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func runeGenerateString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func stringBuilderGenerateString(n int) string {
	b := strings.Builder{}
	for i := 0; i < n; i++ {
		b.WriteByte(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
