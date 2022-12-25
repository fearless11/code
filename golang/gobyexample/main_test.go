package main

import (
	"fmt"
	"testing"
)

// TestIntMinBasic ...
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2,-2) = %d; want -2", ans)
	}
}

func TestIntMinTable(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, t1 := range tests {
		testname := fmt.Sprintf("%d,%d", t1.a, t1.b)
		// t.Run enables running “subtests”
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(t1.a, t1.b)
			if ans != t1.want {
				t.Errorf("got %d, want %d", ans, t1.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
