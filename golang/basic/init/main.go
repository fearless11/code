package main

import (
	"basic/init/aa"
	"fmt"
	"time"
)

// golang中init的顺序
/*
init 顺序
1、在同一个 package 中，可以多个文件中定义 init 方法
2、在同一个 go 文件中，可以重复定义 init 方法
3、在同一个 package 中，不同文件中的 init 方法的执行按照文件名先后执行各个文件中的 init 方法
4、在同一个文件中的多个 init 方法，按照在代码中编写的顺序依次执行不同的 init 方法
5、对于不同的 package，如果不相互依赖的话，按照 main 包中 import 的顺序调用其包中的 init() 函数
6、如果 package 存在依赖，调用顺序为最后被依赖的最先被初始化，例如：导入顺序 main –> A –> B –> C，则初始化顺序为 C –> B –> A –> main，一次执行对应的 init 方法。
所有 init 函数都在同⼀个 goroutine 内执⾏。
所有 init 函数结束后才会执⾏ main.main 函数。
*/

/* output:
aa init 1
aa init 2
main init 1
main init 2
aa hi
main hi
main hello
aa goroutine
aa init goroutine
*/

func init() {
	fmt.Println("main init 1")
}

func init() {
	fmt.Println("main init 2")
}

// Hello 测试hello
func Hello() {
	fmt.Println("main hello")
}

func main() {
	aa.Hi()
	fmt.Println("main hi")
	Hello()
	time.Sleep(1 * time.Second)
}
