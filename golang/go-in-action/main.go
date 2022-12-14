package main

import (
	"log"
	"os"

	_ "go-in-action/matchers"
	"go-in-action/search"
)

// init在main之前调用
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
