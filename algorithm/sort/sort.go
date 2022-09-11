package sort

// 选择排序
// T(n)取决于比较次数
// (n-1)+(n-2)+(n-3)+...+2+1 = n(n-1)/2
// T(n) = O(N^2/2)
// S(n) = 1
// 运行时间于输入无关，数据移动最少，交换N次
func selectSort(a []int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			// compare
			if a[j] < a[min] {
				min = j // get min number index
			}
		}
		// exchange
		a[i], a[min] = a[min], a[i]
	}
	return a
}

// 插入排序
// 应用: 比较有序的数组
// 假设空出首位，比较后把最小的交换移动到首位
// T(n)取决于交换的次数
// 最好情况N-1次比较，0次交换，有序数组
// 最坏情况N^2/2次比较，N^2/2次交换
// 比较 1+2+...+(n-2)+(n-1) = n(n-1)/2
// 交换 1+2+...+(n-2)+(n-1) = n(n-1)/2
// 平均情况N^2/4次比较，N^2/4次交换
func insertSort(a []int) []int {
	n := len(a)
	for i := 1; i < n; i++ {
		// min number exchange head point
		for j := i; j > 0; j-- {
			// compare
			if a[j-1] < a[j] {
				break
			}
			// exchange
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
	return a
}

// 希尔排序
// 排序不稳定
func shellSort(a []int) {

}
