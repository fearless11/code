//Package runner 监控程序的执行时间
//  如果运行时间太长,也可以用runner包终止程序
//  场景： 后台调度任务cron执行,或基于定时任务的云环境
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner任务管理： 添加、启动、执行、额外的中断检测和超时设置

// Runner 在给定的超时时间内执行一组任务，
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	// 发送操作系统的信号
	interrupt chan os.Signal
	// 	任务是否完成
	complete chan error
	// 任务是否超时
	timeout <-chan time.Time

	tasks []func(int)
}

// ErrTimeout 任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 会在接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

// New 工厂函数 返回一个新的准备使用的 Runner
func New(d time.Duration) *Runner {
	return &Runner{
		// 通道至少能接收一个来自语言运行时的 os.Signal 值
		// 确保语言运行时发送这个事件的时候不会被阻塞
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add 将一个任务附加到 Runner 上
// ...可变参数可以接受任意数量的值作为传入参数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	// 我们希望接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的 goroutine 执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务处理完成时发出的信号
	case err := <-r.complete:
		return err
	// 当任务处理程序运行超时时发出的信号
	case <-r.timeout:
		return ErrTimeout
	}
}

// run 执行每一个已注册的任务
func (r *Runner) run() error {

	for id, task := range r.tasks {
		// 检测操作系统的中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// 执行已注册的任务
		// task本身就是传递的函数
		task(id)
	}
	return nil
}

// gotInterrupt 验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件被触发时发出的信号
	case <-r.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
