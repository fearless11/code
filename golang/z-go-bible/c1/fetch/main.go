// Fetch prints the content found at a URL
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func fetch1(in []string) {
	for _, url := range in[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// b, err := ioutil.ReadAll(resp.Body)
		// 防止资源泄露，关闭可读的服务器响应流。
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s:%v\n", url, err)
			os.Exit(1)
		}
		// fmt.Printf("%s", b)
		fmt.Printf("%v\n", resp.StatusCode)
	}
}

//fetchall fetches URLS in parallel and reports their times and sizes
func fetchall(in []string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range in[1:] {
		go fetch(url, ch)
	}
	for range in[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main() {
	/*
		fetch程序所要做的工作基本一致，fetchall的特别之处在于它会同时
		去获取所有的URL，所以这个程序的总执行时间不会超过执行时间最长
		的那一个任务，前面的fetch程序执行时间则是所有任务执行时间之和
	*/
	// fetch1(os.Args)
	fetchall(os.Args)
}
