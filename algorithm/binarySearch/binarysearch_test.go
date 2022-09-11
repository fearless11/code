package search

import (
	"testing"
)

// go test -v -run TestSearch
func TestSearch(t *testing.T) {
	// table test
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{1, 2, 3, 4, 5}, target: 5, want: 4},
		{nums: []int{1, 2, 3, 4, 5}, target: 2, want: 1},
		{nums: []int{1}, target: 1, want: 0},
		{nums: []int{2, 3, 4, 5}, target: 1, want: -1},
	}

	for _, test := range tests {
		if got := search(test.nums, test.target); got != test.want {
			t.Errorf("search(%v, %v), got: %v, want: %v", test.nums, test.target, got, test.want)
		}
	}
}

// go test -v -run TestFirstBadVersion
func TestFirstBadVersion(t *testing.T) {
	var tests = []struct {
		n    int
		want int
	}{
		{n: 1, want: 1},
		{n: 5, want: 1},
	}

	for _, test := range tests {
		got := firstBadVersion(test.n)
		if got != test.want {
			t.Errorf("firstBadVersion(%v) got: %v, want: %v", test.n, got, test.want)
		}
	}
}

// go test -v -run TestSearchInsert
func TestSearchInsert(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, want: 2},
		{nums: []int{1}, target: 1, want: 0},
		{nums: []int{1, 3, 5, 6}, target: 7, want: 4},
	}

	for _, v := range tests {
		got := searchInsert(v.nums, v.target)
		if got != v.want {
			t.Errorf("searchInsert(%v, %v) got: %v,wnat: %v", v.nums, v.target, got, v.want)
		}
	}
}

// go test -v -run TestSearchMatrix
func TestSearchMatrix(t *testing.T) {
	var tests = []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}}, target: 2, want: false},
		{matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}}, target: 3, want: true},
	}

	for _, v := range tests {
		got := searchMatrix(v.matrix, v.target)
		if got != v.want {
			t.Errorf("searchMatrix(%v, %v) got: %v,wnat: %v", v.matrix, v.target, got, v.want)
		}
	}
}
