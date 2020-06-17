package main

import (
	"fmt"
	"os"
	"runtime"
)

type A struct {
	AA int
}

func (a A) add(num A) {
	fmt.Println(a.AA + num.AA)
}

func main() {

	a1 := A{1}
	a2 := A{2}
	way1 := a1.add
	fmt.Printf("%T\n", way1)
	way1(a2)

	way2 := A.add
	fmt.Printf("%T\n", way2)
	way2(a1, a2)

	fmt.Fprintf(os.Stdout, "%T", way2)
	// fmt.Printf("%T", f)
	// defer printStack()

	// 	f(3)

	// 	hello()

	// 	var hello []func()
	// 	for _, v := range []int{1, 2, 3, 4, 5} {
	// 		i := v
	// 		fmt.Println(i)
	// 		hello = append(hello, func() { fmt.Println("hello", i) })
	// 	}

	// 	for _, v := range hello {
	// 		v()
	// 	}
	// }

	// func hello() {
	// 	// 后进先出
	// 	defer func() {
	// 		fmt.Println("aa")
	// 	}()

	// 	defer func() {
	// 		fmt.Println("bb")
	// 	}()

	// 	fmt.Println("hi")
	// }

	// func f(x int) {
	// 	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	// 	defer func() {
	// 		p := recover()
	// 		fmt.Printf("defer %d %v\n", x, p)
	// 	}()
	// 	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
