// 有缓冲的通道（buffered channel）是一种在被接收前能存储一个或者多个值的通道
//   并不强制要求 goroutine 之间必须同时完成发送和接收
//   阻塞情况：
//     通道中没有要接收的值时，接收动作才会阻塞
//     通道没有可用缓冲区容纳被发送的值时，发送动作才会阻塞

////// 使用有缓冲的通道和固定数目的goroutine来处理一堆工作 /////////////
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// go run main.go
/*
worker: 4 started task: 2
worker: 2 started task: 1
worker: 3 started task: 3
worker: 1 started task: 4
worked: 4 completed task: 2
worker: 4 started task: 5
worked: 1 completed task: 4
worker: 1 started task: 6
worked: 3 completed task: 3
worker: 3 started task: 7
worked: 3 completed task: 7
worker: 3 started task: 8
worked: 1 completed task: 6
worker: 1 started task: 9
worked: 2 completed task: 1
worker: 2 started task: 10
worked: 4 completed task: 5
worker: 4 shutting down
worked: 3 completed task: 8
worker: 3 shutting down
worked: 1 completed task: 9
worker: 1 shutting down
worked: 2 completed task: 10
worker: 2 shutting down
*/

const (
	//  4个goroutine并发处理10个任务,处理完退出
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func init() {
	// 初始化随机种子
	rand.Seed(time.Now().Unix())
}

func main() {
	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	// 启动 goroutine 来处理工作
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 产生一组工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("task: %d", post)
	}

	// 当所有工作都处理完时关闭通道
	// 以便所有 goroutine 退出
	close(tasks)

	// // 等待所有工作完成goroutine退出后结束程序
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// 等待分配工作
		task, ok := <-tasks
		if !ok {
			// 这意味着通道已经空了，并且已被关闭
			fmt.Printf("worker: %d shutting down\n", worker)
			return
		}

		fmt.Printf("worker: %d started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("worked: %d completed %s\n", worker, task)
	}
}
