// 这个示例程序展示如何使用 atomic 包里的

////////// Store 和 Load 类函数来提供对数值类型
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	// 给定 goroutine 执行的时间
	time.Sleep(1 * time.Second)

	fmt.Println("shutdown now")
	// 该停止工作了，安全地设置 shutdown 标志
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

// doWork 用来模拟执行工作的 goroutine，
// 检测之前的 shutdown 标志来决定是否提前终止
func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)

		// 要停止工作了吗？
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("shutting %s down\n", name)
			break
		}
	}
}
