// 用atmoic包提供对数值类型的安全访问

/////  atomic.AddInt64原子操作安全累加
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 安全对counter加1
		// 强制同一时刻只能有一个goroutine运行并完成这个加法操作
		atomic.AddInt64(&counter, 1)
		// 当前goroutine从线程中退出，并放回队列
		runtime.Gosched()
	}
}
