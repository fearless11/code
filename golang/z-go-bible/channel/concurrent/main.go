// 并发循环

package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	filenames := []string{"1.txt", "2.txt", "3.txt"}
	// printer1(filenames)
	fmt.Println("hi2: ", printer2(filenames))
	fmt.Println("hi3: ", printer3(filenames))
	fmt.Println("hi good printer:", goodPrinter(filenames))
}

//printer1 通过共享channel通知外部的goroutine
func printer1(filenames []string) {

	ch := make(chan struct{})
	for _, f := range filenames {
		// 匿名函数中的循环变量快照问题
		// 打印的为slice中最后一个元素
		// go func() {
		// 	log.Println(f)
		// 	ch <- struct{}{}
		// }()
		go func(f string) {
			log.Println(f)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}

// printer2 微妙的bug,当i等于1时返回到调用方，
// 没有一个goroutine去排空errors channel。
// 剩下的worker goroutine在向这个channel中发送值时，
// 都会永远地阻塞下去，并且永远都不会退出。这种情况叫做goroutine泄露。
// 可能会导致整个程序卡住或者跑出out of memory的错误。
func printer2(filenames []string) int {
	num := make(chan int)

	for i, f := range filenames {
		go func(i int, f string) {
			log.Println("go2:", f)
			num <- i
		}(i, f)
	}

	for range filenames {
		i := <-num
		log.Println("go2:", i)
		if i == 1 {
			return i // 返回时可能存在goroutine泄露
		}
	}

	return 99
}

// printer3 解决goroutine泄漏
// 用一个具有合适大小的buffered channel,
// 这样worker goroutine向channel中发送错误时就不会被阻塞。
// 另外启动一个goroutine来排空channel也可以。
func printer3(filenames []string) int {
	num := make(chan int, len(filenames))

	for i, f := range filenames {
		go func(i int, f string) {
			log.Println("go3:", f)
			num <- i
		}(i, f)
	}

	for range filenames {
		i := <-num // channel会阻塞直到没有数据
		log.Println("go3:", i)
		if i == 1 {
			return i
		}
	}
	return 99
}

// goodPrinter
// 计数器实现goroutine什么时候结束
// goroutine启动时加一，退出时减一。特殊计数器并发安全
func goodPrinter(filename []string) int {
	counter := make(chan int)
	var wg sync.WaitGroup
	for _, f := range filename {
		// Add是为计数器加一，必须在worker goroutine开始之前调用
		wg.Add(1)
		// work
		go func(f string) {
			defer wg.Done()
			counter <- 1
		}(f)
	}

	go func() {
		wg.Wait()
		close(counter)
	}()

	var total int
	// channel只要有数据阻塞
	for i := range counter {
		total += i
	}

	return total
}
