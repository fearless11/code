package main

import "fmt"

func main() {
	// a()
	// fmt.Println(b())
	// c()
	d()
}

//////////////////// defer //////////////////////
// output: 0
// defer 参数是在被defer语句包裹的时候进行求值
func a() {
	i := 0
	// 在defer运行时, i被求值为0, 传入Print函数
	defer fmt.Println(i)
	i++
	return
}

// output: 2
// 多个defer语句,遵循FILO后进先出原则
func b() (i int) {
	defer func() { i++ }()
	return 1
}

// output: hello world
// defer 可读取函数返回值,进行操作,甚至修改返回值
func c() {
	defer fmt.Println("world")
	defer fmt.Println("hello")
	return
}

////////////////////// panic  //////////////////////
/*
Panic 语句可以让程序陷入一种恐慌的状态，程序会停止原先的 workflow，然后按顺序依次执行 defer 过的语句，
然后返回调用的函数，对于调用者而言，此时也陷入了恐慌的状态。然后持续之前的动作，直到此时的 goroutine 中
全部返回，然后程序 crash 掉。

整个程序的 panic 状态可能有多个原因造成，数组的越界或者对 nil 的操作都会导致程序的 panic 状态,
显示的调用 panic() 函数也会导致程序进入 panic 状态。

goroutine & painc
Go 语言中，defer 调用的时候，会将调用的 goroutine 的协程和 defer 调用的函数进行关联。
defer 调用的函数对应的各自协程中的状态。所以，如果遇到了跨协程的时候，跨协程的 defer recover 会失效，即两个协程之间互不干扰。
一个 goroutine 在 panic 的时候，不会调用其他 goroutine 的延迟函数。
但是一个 goroutine 的运行的panic，如果没有很好的处理，也会影响其他 goroutine 的运行。
所以，对于多协程的时候，应该各自的 goroutine 处理各自的 panic 状况。
*/

///////////////////// recover ////////////////////////
/*
Recover 语句则是用来对 panic 状态的程序进行恢复的语句。recover 函数只有其在 defer 中进行调用的时候才会
有效，如果在 defer 外进行调用，则无效。如果程序没有进入恐慌状态，那么 recover 函数则返回 nil，反之，
recover 函数返回调用 panic 函数时的参数。然后让程序恢复正常。
*/

/* output:
calling e
printing in e 0
printing in e 1
printing in e 2
panicing!
defer in e 2
defer in e 1
defer in e 0
recovered in d 3
*/
func d() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in d", r)
		}
	}()
	fmt.Println("calling e")
	e(0)
	fmt.Println("returned normally from e")
}

func e(i int) {
	if i > 2 {
		fmt.Println("panicing!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("defer in e", i)
	fmt.Println("printing in e", i)
	e(i + 1)
}
