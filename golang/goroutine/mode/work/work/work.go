//Package work 如何使用无缓冲通道创建一个goroutine池，并发控制一组工作，让其并发执行
//  无缓冲优点: 不会有工作在队列里丢失或卡主，所有工作会被处理
package work

import "sync"

////////// 动态决定创建多少goroutine个数来处理任务

// Worker 必须满足接口类型
type Worker interface {
	Task()
}

// Pool 提供一个goroutine池
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 创建一个工作池
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run 提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown 等待所有goroutine停止工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
