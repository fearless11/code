package sort

import (
	"leetcode/util"
	"testing"
)

// go test -v -count 1 -run TestSelectSort
func TestSelectSort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{5, 3, 2, 1, 7},
			want: []int{1, 2, 3, 5, 7},
		},
		{
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
	}

	for _, v := range tests {
		got := selectSort(v.nums)
		if !util.IsEqual(got, v.want) {
			t.Errorf("selectSort(%v) got: %v,wnat: %v", v.nums, got, v.want)
		}
	}
}

// go test -v -count 1 -run TestInsertSort
func TestInsertSort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{5, 3, 2, 1, 7},
			want: []int{1, 2, 3, 5, 7},
		},
		{
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
	}

	for _, v := range tests {
		got := insertSort(v.nums)
		if !util.IsEqual(got, v.want) {
			t.Errorf("insertSort(%v) got: %v,wnat: %v", v.nums, got, v.want)
		}
	}
}

// go test -bench=BenchmarkSelectSort -count=3 -benchmem
func BenchmarkSelectSort(b *testing.B) {
	a := util.RandomIntArray(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		selectSort(a)
	}
}

// go test -bench=BenchmarkInsertSort -count=3 -benchmem
func BenchmarkInsertSort(b *testing.B) {
	a := util.RandomIntArray(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		insertSort(a)
	}
}
