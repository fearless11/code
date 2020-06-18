// dup prints the text of echo line that appears to
// more than once in the standard input, preceded(优于)
// by its count
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
func dup2(in []string) {
	counts := make(map[string]int)
	files := in[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()

			// 修改dup2，出现重复的行时打印文件名称
			for line, n := range counts {
				if n > 1 {
					fmt.Println(arg, n, line)
				}
			}
		}
	}

	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Println(n, line)
	// 	}
	// }
}

func dup3(in []string) {
	counts := make(map[string]int)
	for _, filename := range in[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
}
func main() {
	// 以"流”模式读取输入，并根据需要拆分成多个行。
	// 理论上，这些程序可以处理任意数量的输入数据
	dup1()
	// dup2(os.Args)

	//一口气把全部输入数据读到内存中，一次分割为多行，然后处理它们
	// dup3(os.Args)
}
