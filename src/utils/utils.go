package utils

import (
	"hash/fnv"
	"math/rand"
)

// NonEmpty checks whether a string is empty
func NonEmpty(s string) bool {
	return len(s) != 0
}

//Hash return of a string
func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes random number of given bit
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

