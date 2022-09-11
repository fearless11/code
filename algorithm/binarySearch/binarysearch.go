package search

import (
	"math/rand"
	"time"
)

/* binary search
1. set middle pointer: left + (right-left)/2
2. juge target, move left and right pointer: left = mid + 1; right = mid
3. judge other condition
*/

/* T(n) = O(logN)
第一次区间：N/1
第二次区间：N/2
第三次区间：N/2/2
第四次区间：N/2/2/2
第五次区间：N/2^(5-1)
...
第k次区间：N/2^(k-1)

最坏的情况区间为1才找到：
1 = N/2^(k-1)
2^(k-1) = N
k -1  = log2(N)
logN = (k-1) * log2
k -> logN
*/

// https://leetcode-cn.com/problems/binary-search
// 有序数组和目标，目标存在返回下标，否则返回-1
// input: nums = [-1,0,3,5,9,12], target = 9
// output: 4
// input: nums = [-1,0,3], target = 2
// output: -1
// T(n) = O(logN); S(n) = O(1)
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == right {
		if nums[left] == target {
			return left
		}
	}

	return -1
}

// https://leetcode-cn.com/problems/first-bad-version/
// 有 n 个版本 [1, 2, ..., n]，找出导致之后所有版本出错的第一个错误的版本
// input: n = 5, bad = 4
// output: 4
// reason:
// isBadVersion(3) -> false
// isBadVersion(5) -> true
// isBadVersion(4) -> true
func firstBadVersion(n int) int {
	left := 1
	right := n
	for left < right {
		mid := left + (right-left)/2
		if !isBadVersion(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func isBadVersion(n int) bool {
	num := n + 10
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(num) + 1
	if r == n {
		return true
	}
	return false
}

// https://leetcode-cn.com/problems/search-insert-position/
// 有序数组和目标值，找到目标则返回其索引，否则返回案顺序插入的位置索引
// input: nums = [1,3,5,6], target = 5
// output: 2
// input: nums = [1,3,5,6], target = 2
// output: 1
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	tag := left

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			tag = mid
			break
		}

		if nums[mid] < target {
			left = mid + 1
			tag = left
		} else {
			right = mid
		}
	}

	if target > nums[right] {
		tag = right + 1
	}

	return tag
}

// https://leetcode-cn.com/problems/search-a-2d-matrix/submissions
func searchMatrix(matrix [][]int, target int) bool {
	nums := []int{}
	lenout := len(matrix)
	for i := 0; i < lenout; i++ {
		len := len(matrix[i])
		for j := 0; j < len; j++ {
			nums = append(nums, matrix[i][j])
		}
	}
	return searchTrueFalse(nums, target)
}

func searchTrueFalse(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == right {
		if nums[left] == target {
			return true
		}
	}

	return false
}
