package main

import (
	"code/basic/init/aa"
	"fmt"
	"time"
)

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
