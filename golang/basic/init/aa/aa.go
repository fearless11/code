// Package aa 包提供了测试init函数的功能
package aa

import "fmt"

func init() {
	fmt.Println("aa init 1")
}

func init() {
	fmt.Println("aa init 2")
}

func init() {
	go func() {
		fmt.Println("aa init goroutine")
	}()
}

// Hi say hi
func Hi() {
	fmt.Println("aa hi")

	go func() {
		fmt.Println("aa goroutine")
	}()
}
