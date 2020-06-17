// 互斥锁 mutex
//   同步访问共享资源的方式
//   互斥锁用于在代码上创建一个临界区，保证同一时间只有一个 goroutine 可以执行这个临界区代码

///// 临界区累加
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go run main.go
// output:
/*
id 2 counter 1
id 1 counter 2
id 2 counter 3
id 1 counter 4
Final Counter: 4
*/

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {

	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func incCounter(id int) {
	defer wg.Done()

	// 执行四次，1 + 1 + 1 + 1 = 4
	for count := 0; count < 2; count++ {

		// mutex的必须等一个goroutine操作完成
		mutex.Lock()
		{
			value := counter
			// 当前goroutine从线程退出，并放回队列
			runtime.Gosched()
			value++
			counter = value
			fmt.Printf("id %v counter %v\n", id, counter)
		}
		mutex.Unlock()

	}
}
