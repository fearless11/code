package doublepointer

import (
	"leetcode/util"
	"testing"
)

// go test -v -run TestSortedSquares
func TestSortedSquares(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{nums: []int{-7, -3, 2, 3, 11}, want: []int{4, 9, 9, 49, 121}},
		{nums: []int{-7, -6, -5, -4}, want: []int{16, 25, 36, 49}},
	}

	for _, v := range tests {
		got := sortedSquares(v.nums)
		if !util.IsEqual(got, v.want) {
			t.Errorf("sortedSquares(%v) got: %v,wnat: %v", v.nums, got, v.want)
		}
	}
}

// go test -v -run TestRotate
func TestRotate(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
		want []int
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3, want: []int{5, 6, 7, 1, 2, 3, 4}},
	}

	for _, v := range tests {
		rotate(v.nums, v.k)
		got := v.nums
		if !util.IsEqual(got, v.want) {
			t.Errorf("rotate(%v,%v) got: %v,wnat: %v", v.nums, v.k, got, v.want)
		}
	}
}

// go test -v -run TestMoveZeroes
func TestMoveZeroes(t *testing.T) {
	var tests = []struct {
		nums []int
		want []int
	}{
		{nums: []int{1, 0}, want: []int{1, 0}},
		{nums: []int{2, 1}, want: []int{2, 1}},
		{nums: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
	}

	for _, v := range tests {
		moveZeroes(v.nums)
		got := v.nums
		if !util.IsEqual(got, v.want) {
			t.Errorf("moveZeroes(%v) got: %v,wnat: %v", v.nums, got, v.want)
		}
	}
}

// go test -v -run TestTwoSum
func TestTwoSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{1, 2}},
		{nums: []int{2, 3, 4}, target: 6, want: []int{1, 3}},
		{nums: []int{-1, 0}, target: -1, want: []int{1, 2}},
	}

	for _, v := range tests {
		got := twoSum(v.nums, v.target)
		if !util.IsEqual(got, v.want) {
			t.Errorf("twoSum(%v,%v) got: %v,wnat: %v", v.nums, v.target, got, v.want)
		}
	}
}
