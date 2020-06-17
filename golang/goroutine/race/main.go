// 竞争状态
//   多个goroutine在没有互相同步的情况下，访问某个共享资源，
//   对共享资源的操作必须原子化
package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
竞争检测器标志来编译程序
go build -race
运行程序，如果有竞争会显示出来
./race
*/

/*
可以明确看到是哪里代码有冲突，哪个goroutine引发数据竞争
output
==================
WARNING: DATA RACE
Read at 0x00000121d848 by goroutine 7:
  main.incCounter()
      code/goroutine/race/main.go:31 +0x6f

Previous write at 0x00000121d848 by goroutine 6:
  main.incCounter()
      code/goroutine/race/main.go:37 +0x90

Goroutine 7 (running) created at:
  main.main()
       code/goroutine/race/main.go:21 +0x89

Goroutine 6 (finished) created at:
  main.main()
       code/goroutine/race/main.go:20 +0x68
==================
Final Counter: 4
Found 1 data race(s)
*/

var (
	counter int
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
		value := counter

		// 当前goroutine从线程中退出，并放回队列
		runtime.Gosched()

		value++
		// 并发时竞争点： value:=counter & counter=value
		// 并发时一个goroutine的操作value++还没有赋值给counter，
		// 另一个goroutine重新赋值给value为之前的counter
		counter = value

		fmt.Printf("id %v counter %v\n", id, counter)
	}

}
