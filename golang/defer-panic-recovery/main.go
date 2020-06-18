package main

import (
	"fmt"
	"time"
)

/*
defer: 延迟处理
panic: 中断程序
recovery: 截取panic恢复程序

同一个goroutine中可用recovery恢复panic
*/

func main() {
	// deferParam()
	// fmt.Println(deferIadd())
	// deferFILO()
	// recoveryPanic()
	panicBygoroutine()
}

//////////////////// defer //////////////////////
// output: 0
// defer参数是在被defer语句包裹的时候进行求值
// 在defer运行时i被求值为0传入Print函数
func deferParam() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// output: 2
// defer读取函数返回值进行操作,甚至修改返回值
func deferIadd() (i int) {
	defer func() { i++ }()
	return 1
}

// output: hello world
// 多个defer遵循FILO后进先出原则
func deferFILO() {
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
calling autoPanic
printing in autoPanic 0
printing in autoPanic 1
printing in autoPanic 2
auto panicing!
defer in autoPanic 2
defer in autoPanic 1
defer in autoPanic 0
recovered in recoveryPanic 3
*/
// defer的后进先出
// 在同一个goroutine的defer中用recovery捕获panic
func recoveryPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in recoveryPanic", r)
		}
	}()
	fmt.Println("calling autoPanic")
	autoPanic(0)
	fmt.Println("returned normally from autoPanic")
	select {}
}

func autoPanic(i int) {
	if i > 2 {
		fmt.Println("auto panicing!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("defer in autoPanic", i)
	fmt.Println("printing in autoPanic", i)
	autoPanic(i + 1)
}

/*
output:
defer func here
defer goroutine
panic: auto panic
*/
// panic只管当前goroutine中的defer会被调用到，不保证其他defer调用
func panicBygoroutine() {
	defer fmt.Println("defer main") // 不会执行
	go func() {
		defer fmt.Println("defer goroutine")
		func() {
			defer func() {
				fmt.Println("defer func here")
			}()
			panic("auto panic")
		}()
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("main ending")
}
