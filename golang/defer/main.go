package main

import "fmt"

func deferTest1() {
	i := 1
	defer fmt.Println(i)
	i = 3
}

func deferTest2() {
	i := 1
	defer func() {
		fmt.Println(i)
	}()
	i = 3
}

func deferTest3() {
	{
		defer fmt.Println("defer code block 1")
		fmt.Println("this is code block 1")
	}
	defer fmt.Println("defer in main func")
}

func main() {
	deferTest1()
	deferTest2()
	deferTest3()
}
