//Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

// problem
/*
每次循环迭代字符串s的内容都会更新。+=连接原字符串、空格和下个参数，
产生新字符串, 并把它赋值给s。s原来的内容已经不再使用，
将在适当时机对它进行垃圾回收。
*/
func echo1(in []string) {
	var s, seq string
	for i := 1; i < len(os.Args); i++ {
		s += seq + in[i]
		seq = " "
	}
	fmt.Println(s)
}

func echo2(in []string) {
	s, seq := "", ""
	for _, arg := range in[1:] {
		s += seq + arg
		seq = " "
	}
	fmt.Println(s)
}

func main() {
	// 低效的版本
	echo1(os.Args)
	echo2(os.Args)
	// 简单且高效的解决方案
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args)
}
