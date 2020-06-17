// goroutine调度器如何在单个线程上切分时间片

//////// 求5000内的素数
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go run main.go
/*
create goroutines
waiting to finish
B:2
B:3
...
A:431         切换goroutine
A:433
A:439
B:479         切换goroutine
B:487
B:491
completed B
A:479         切换goroutine
A:487
completed A
*/

var wg sync.WaitGroup

func main() {
	// 修改任何语言运行时配置参数的时候，需要配合基准测试来评估程序运行结果
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	fmt.Println("create goroutines")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("waiting to finish")
	wg.Wait()

	fmt.Println("terminating program")
}

// 显示5000内的素数
func printPrime(perfix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}

		fmt.Printf("%s:%d\n", perfix, outer)
	}

	fmt.Println("completed", perfix)

}
