package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)

	// 非并发遍历
	// go func() {
	// 	for _, root := range roots {
	// 		walkDir1(root, fileSizes)
	// 	}
	// 	close(fileSizes)
	// }()

	// 并发遍历目录
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// 通过开关控制-v确认是否显示进度, 开启则500毫秒显示一次
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nsizes int64

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				// chan关闭,通过标签跳出select和for循环
				// 简单break只会跳出select循环
				break loop
			}
			nfiles++
			nsizes += size
		case <-tick:
			printDiskUsage(nfiles, nsizes)
		}
	}
	printDiskUsage(nfiles, nsizes)

}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1fGB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir1 非并发
func walkDir1(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir1(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// walkDir 递归可以实现遍历,并发
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// 信号量控制并发
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // 获得信号量
	defer func() { <-sema }() // 释放信号量

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v", err)
		return nil
	}
	return entries
}
