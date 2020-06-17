//多个goroutine通过channel相连形成pipeline管道
// 单个方向channel
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squarers := make(chan int)
	go counter(naturals)
	go squarer(squarers, naturals)
	printer(squarers)
}

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
