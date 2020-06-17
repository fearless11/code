package main

import (
	"context"
	"fmt"
	"time"
)

// 自定义键的类型，避免与作用域内置类型发生碰撞
type key string

var userKey = key("test")

func main() {
	// withCancel()
	// withDeadline()
	withValue()
}

// 通知防止goroutine泄露
func withCancel() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1

		//  子goroutine
		go func() {
			for {
				select {
				case <-ctx.Done():
					// returning保证没有goroutine泄露
					return
				case dst <- n:
					n++
				}
			}
		}()

		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	// 返回时发送cancel()的取消通知
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

// 设置任务的执行时间
func withDeadline() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	// do somthing on deadline

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

// 传递值
func withValue() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, userKey, "hi")
	defer cancel()
	a(ctx)
}

func a(ctx context.Context) {
	fmt.Println(ctx.Value(userKey))
	// 修改传递的值
	ctx = context.WithValue(ctx, userKey, "world")
	b(ctx)
}

func b(ctx context.Context) {
	fmt.Println(ctx.Value(userKey))
}
