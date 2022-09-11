package doublepointer

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/
func sortedSquares(nums []int) []int {
	// diff left and right, get max to tag, then move left and right and tag
	// T(n)=O(n)
	// S(n)=O(n)

	n := len(nums)
	result := make([]int, n)
	i := 0
	j := n - 1
	tag := n - 1

	for i <= j {
		left := nums[i] * nums[i]
		right := nums[j] * nums[j]
		if left > right {
			result[tag] = left
			i++
		} else {
			result[tag] = right
			j--
		}
		tag--
	}
	return result
}

// https://leetcode-cn.com/problems/rotate-array/
func rotate(nums []int, k int) {
	// 输入: nums = [1,2,3,4,5,6,7], k = 3
	// 输出: [5,6,7,1,2,3,4]
	// 解释:
	// 向右轮转 1 步: [7,1,2,3,4,5,6]
	// 向右轮转 2 步: [6,7,1,2,3,4,5]
	// 向右轮转 3 步: [5,6,7,1,2,3,4]
	rotate2(nums, k)
}

// T(n)=O(n)  S(n)=O(n)
func rotate1(nums []int, k int) {
	n := len(nums)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		idx := (i + k) % n
		temp[idx] = nums[i]
	}
	copy(nums, temp)
}

// T(n)=O(n) S(n)=O(1)
func rotate2(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// https://leetcode-cn.com/problems/move-zeroes/
// T(n)=O(n) S(n)=O(1)
func moveZeroes(nums []int) {
	// 逆向思路：非0时考虑交换，把所有非0的移动到左边，剩下右边都是0
	// 左右指针都指向开始点，右指针一直移动找非0，非0同时移动，为0是左指针不动，等待右指针找到非0的交换
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
func twoSum(numbers []int, target int) []int {
	return twoSum2(numbers, target)
}

// T(n)=O(n^2) S(n)=O(1)
func twoSum1(numbers []int, target int) []int {
	n := len(numbers)
	index1, index2 := 0, 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := numbers[i] + numbers[j]
			if sum == target {
				index1 = i
				index2 = j
			}
		}
	}
	return []int{index1 + 1, index2 + 1}
}

// T(n)=O(n) S(n)=O(1)
func twoSum2(numbers []int, target int) []int {
	n := len(numbers)
	index1, index2 := 0, n-1
	for index1 <= index2 {
		sum := numbers[index1] + numbers[index2]
		if sum > target {
			index2--
		} else if sum < target {
			index1++
		} else {
			break
		}
	}
	return []int{index1 + 1, index2 + 1}
}
