package main

import (
	"log"
	"os"

	_ "goinaction/matchers"
	"goinaction/search"
)

// init在main之前调用
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

func main() {
	// 搜索president总统关键字的rss
	search.Run("president")
}
