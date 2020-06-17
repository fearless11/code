package main

import "fmt"

/*
执行时指定GOMAXPROCS查看结果不一样
GOMAXPROCS=1 go run main.go
输出：111111100000111111100000

GOMAXPROCS=2 go run main.go
输出：101010101011010101

因为指定GOMAXPROCS为1时,只能运行一个goroutine,只有等待该goroutine执行一段时间间休眠后,才能执行调度其他goroutine执行
M:N 此时n为1

当GOMAXPROCS为2时,可以运行2个goroutine,故交替打印

*/
func main() {
	for {
		go fmt.Print("0")
		fmt.Print("1")
	}
}
