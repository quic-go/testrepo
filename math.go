package testrepo

import "pgregory.net/rand"

func Add(a, b int) int {
	return a + b
}

func Rand() int {
	return rand.Int()
}
