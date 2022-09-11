package util

import "testing"

// go test -run TestRandomIntArray -v
func TestRandomIntArray(t *testing.T) {
	t.Log(RandomIntArray(10))
}
