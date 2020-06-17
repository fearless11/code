// 有缓冲通道使用示例
// 通过数据共享消息，而不是共享数据
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// wg等待程序完成
var wg sync.WaitGroup

func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 任务通道为10
	tasks := make(chan string, 10)

	// 创建4个goroutine处理10个工作
	for gr := 1; gr <= 4; gr++ {
		go worker(tasks, gr)
	}

	// 创建10个工作
	for post := 1; post <= 10; post++ {
		tasks <- fmt.Sprintf("Task： %d", post)
	}

	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		// 接收工作, 当channal关闭时return
		task, ok := <-tasks
		if !ok {
			fmt.Printf("worker: %d: shutting Down\n", worker)
			return
		}
		// 开始工作
		fmt.Printf("worker: %d : started %s\n", worker, task)
		// 模拟在处理工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		// 完成工作
		fmt.Printf("worker: %d : end %s\n", worker, task)
	}
}

// output：
// worker: 4 : started Task： 4
// worker: 1 : started Task： 2
// worker: 2 : started Task： 1
// worker: 3 : started Task： 3
// worker: 4 : end Task： 4
// worker: 4 : started Task： 5
// worker: 2 : end Task： 1
// worker: 2 : started Task： 6
// worker: 4 : end Task： 5
// worker: 4 : started Task： 7
// worker: 2 : end Task： 6
// worker: 2 : started Task： 8
// worker: 1 : end Task： 2
// worker: 1 : started Task： 9
// worker: 3 : end Task： 3
// worker: 3 : started Task： 10
// worker: 1 : end Task： 9
// worker: 1: shutting Down
// worker: 2 : end Task： 8
// worker: 2: shutting Down
// worker: 4 : end Task： 7
// worker: 4: shutting Down
// worker: 3 : end Task： 10
// worker: 3: shutting Down
