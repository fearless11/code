package util

import (
	"math/rand"
)

func IsEqual(a, b []int) bool {
	lena := len(a)
	lenb := len(b)
	if lena != lenb {
		return false
	}

	for i := 0; i < lena; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func RandomIntArray(n int) []int {
	a := []int{}
	for i := 0; i < n; i++ {
		r := rand.Intn(n)
		a = append(a, r)
	}
	return a
}
