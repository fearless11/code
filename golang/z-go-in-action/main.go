package main

import (
	"log"
	"os"

	_ "gitee.com/feareless11/book/code/go-practice/matchers"
	"gitee.com/feareless11/book/code/go-practice/search"
)

// init在main之前调用
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	// 搜索president总统关键字的rss
	search.Run("president")
}
