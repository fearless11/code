package main

import "fmt"

func main() {

	// 递归
	n := 5
	fmt.Printf("%v! = %v\n", n, factorial(n))
	fmt.Printf("%v! = %v\n", n, factorialTail(n, 1))

	// 查找
	searchLinear()
	searchBinary()

	// 排序
	sortBubble()
	sortSelect()
	array := []int{5, 3, 2, 1, 4}
	sortQuick(array, 0, 4)
	fmt.Println("quick sort", array)

	// 去重
	hasDuplicateValue1()
	hasDuplicateValue2()
}

// 阶乘
/*
 1! = 1
 2! = 1 * 2
 3! = 1 * 2 * 3
 4! = 1 * 2 * 3 * 4
 5! = 1 * 2 * 3 * 4 * 5
 n! = n * f(n-1)
 基准情形： n = 1 跳出
*/
// 递归
// 空间复杂度: O(n)
// 问题：栈溢出 100！= 0
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 尾递归
// 问题： 100！= 0
// 递归栈溢出原因：需要保留每一步的值，n个数则需要超过n个值
// 尾递归：只存在一个调用栈，保留的每一步的值用参数存取
func factorialTail(n int, total int) int {
	if n == 1 {
		return total
	}

	return factorialTail(n-1, n*total)
}

// 线性查找
// 时间复杂度：O(N)
func searchLinear() {
	num := 7
	numbers := [5]int{1, 4, 5, 6, 7}

	for i := 0; i < len(numbers); i++ {
		if num == numbers[i] {
			fmt.Printf("the position numbers[%v] find the number %v\n", i, num)
			return
		}
	}
}

// 二分查找
// 时间复杂度：O(logN)
func searchBinary() {
	num := 5
	numbers := [5]int{1, 3, 4, 5, 7}

	// 上下边界
	lowerBound := 0
	upperBound := len(numbers) - 1
	for lowerBound <= upperBound {
		// 取中间元素
		midpoint := (lowerBound + upperBound) / 2
		midvalue := numbers[midpoint]
		if num < midvalue {
			upperBound = midpoint - 1
		} else if num > midvalue {
			lowerBound = midpoint + 1
		} else if num == midvalue {
			fmt.Printf("the position numbers[%v] find the number %v\n", midpoint, midvalue)
			return
		}
	}
}

// 冒泡排序
// 时间复杂度：O(N^2)
// 操作步数分析：比较 & 交换
// 最坏情况：一个长度为5的逆序数组，比较4+3+2+1=10次，交换4+3+2+1=10次。步数为20次，约为5*5=25次
func sortBubble() {
	numbers := []int{3, 6, 7, 2, 4}
	unsortedIndex := len(numbers) - 1
	// 一个轮回中没有发生任何交换则数组已经排序完成
	sorted := false

	for sorted != true {
		sorted = true
		// 轮回
		for i := 0; i < unsortedIndex; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				sorted = false
			}
		}
		unsortedIndex = unsortedIndex - 1
	}

	fmt.Println("bubble sort", numbers)
}

// 选择排序
// 比较冒泡，少了大量交换。 N^2/2
// 时间复杂度：O(N^2)
func sortSelect() {
	numbers := []int{3, 6, 7, 2, 4}

	for i := 0; i < len(numbers); i++ {
		// 默认每次起点值为最小值
		var lowerNumberIndex = i
		// 比较：起点后面的值与最小值比较
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[lowerNumberIndex] {
				lowerNumberIndex = j
			}
		}

		// 交换：最小值与本次检测起点交换
		if lowerNumberIndex != i {
			var temp = numbers[i]
			numbers[i] = numbers[lowerNumberIndex]
			numbers[lowerNumberIndex] = temp
		}
	}

	fmt.Println("select sort", numbers)
}

// 插入排序
// 时间复杂度: O(N^2)
func sortInsert() {
	numbers := []int{3, 6, 7, 2, 4}

	for i := 1; i < len(numbers); i++ {
		// 一次轮回
		// 移出
		position := i
		temp := numbers[i]
		// 比较 : 与左边值比较，左边大则左边数据右移，左边小则当前位置插入临时值，或者左边没有数据插入临时值。
		for position > 0 && numbers[position-1] > temp {
			// 平移
			numbers[position] = numbers[position-1]
			position = position - 1
		}
		// 插入
		numbers[position] = temp
	}
}

// 快排：快速排序
// 时间复杂度：
//   最好 O(NlogN) 平均 O(NlogN)  最坏 O(N^2)
func sortQuick(array []int, leftIndex int, rightIndex int) {
	// 基准情形
	if rightIndex <= leftIndex {
		return
	}
	// 分而治之
	// 一次分区后轴的位置
	pivotPosition := sortPartition(array, leftIndex, rightIndex)
	// 左侧
	sortQuick(array, leftIndex, pivotPosition-1)
	// 右侧
	sortQuick(array, pivotPosition+1, rightIndex)
}

// 分区
func sortPartition(array []int, leftPointer int, rightPointer int) int {
	pivotPosition := rightPointer
	pivot := array[pivotPosition]

	rightPointer = rightPointer - 1
	for {
		for array[leftPointer] < pivot {
			leftPointer = leftPointer + 1
		}

		for array[rightPointer] > pivot {
			rightPointer = rightPointer - 1
		}

		if leftPointer > rightPointer {
			break
		} else {
			// 交换左右指针的值
			array[leftPointer], array[rightPointer] = array[rightPointer], array[leftPointer]
		}
	}

	// 值放到合适位置
	array[leftPointer], array[pivotPosition] = array[pivotPosition], array[leftPointer]
	return leftPointer
}

// 数组中是否有重复的值
// 时间复杂度：O(N^2)
func hasDuplicateValue1() {
	array := []int{1, 3, 5, 7, 9, 3}

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			// 自身不与自身比较
			if i != j && array[i] == array[j] {
				fmt.Println("has duplicate value", array[i])
				return
			}
		}
	}
}

// 利用散列表的key检查存在
// 时间复杂度为 O(N)
// 空间复杂度为 O(N)
func hasDuplicateValue2() {
	array := []int{1, 3, 5, 7, 9, 3}
	// make初始化，返回类型为T
	// new分配置零的存储，返回类型为*T
	mapArray := make(map[int]bool)

	for i := 0; i < len(array); i++ {
		if _, exist := (mapArray[array[i]]); !exist {
			mapArray[array[i]] = true
		} else {
			fmt.Println("has duplicate value", array[i])
		}
	}
}
