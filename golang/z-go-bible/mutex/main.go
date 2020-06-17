// 无阻塞缓存

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"gitee.com/fearless11/explore/go-bible/mutex/memo1"
	"gitee.com/fearless11/explore/go-bible/mutex/memochannel"
	"gitee.com/fearless11/explore/go-bible/mutex/memoshare"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

var incomingURLs = []string{"http://translate.google.cn", "http://www.baidu.com", "http://www.baidu.com", "http://translate.google.cn", "http://www.baidu.com"}

// 顺序执行
func orderMemo1() {
	m := memo1.New(httpGetBody)
	for _, url := range incomingURLs {
		start := time.Now()
		values, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s %s, %d\n", url, time.Since(start), len(values.([]byte)))
	}
}

// 并发执行,不安全
func concurrentMemo1() {
	m := memo1.New(httpGetBody)
	var wg sync.WaitGroup
	for _, url := range incomingURLs {

		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			values, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s %s, %d\n", url, time.Since(start), len(values.([]byte)))
		}(url)
	}

	wg.Wait()
}

// 共享变量方式实现
// 并发 、不重复、无阻塞
func concurrentMemoByShareVar() {
	m := memoshare.New(httpGetBody)
	var wg sync.WaitGroup
	for _, url := range incomingURLs {

		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			values, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s %s, %d\n", url, time.Since(start), len(values.([]byte)))
		}(url)
	}

	wg.Wait()
}

// 通信方式实现
// 并发、不重复、无阻塞
func concurrentMemoByChannel() {
	m := memochannel.New(httpGetBody)
	var wg sync.WaitGroup
	for _, url := range incomingURLs {

		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			values, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s %s, %d\n", url, time.Since(start), len(values.([]byte)))
		}(url)
	}
	wg.Wait()
	m.Close() // 防止channel泄露,内存泄露
}

func main() {
	// orderMemo1()
	// concurrentMemo1()
	// concurrentMemoByShareVar()
	concurrentMemoByChannel()
}
