package generator

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz-1234567890")


// Path return a new random path for url.
func Path() string {
	var generatedURL string

	generatedURL = randSeq(8)
	return generatedURL
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}