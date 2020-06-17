package main

import (
	"log"
	"sync"
	"time"

	"code/goroutine/mode/work/work"
)

// 一段通过goroutine并发向通道提交任务
// 另一端通过goroutine从通道取走任务
// 因为是无缓冲的通道，任务不存在丢失的情况

// names 提供了一组用来显示的名字
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	// 生产：实现了500个goroutine在生产任务（每个goroutine生产一个）
	// 消费：通过work池的2个goroutine在消费任务

	//创建包含两个goroutine的工作池
	p := work.New(2)

	var wg sync.WaitGroup

	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				// 将任务提交执行。当Run返回时，我们就知道任务已经处理完成
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// 让工作池停止工作，等待所有现有的工作完成
	p.Shutdown()
}
