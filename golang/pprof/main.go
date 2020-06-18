package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	pprofNetHTTP()
}

// web形式：只用导入_net/http/pprof后http.ListenAndServe(":6060", nil)即可

// http://localhost:6060/debug/pprof/
// 网页查看观察应用程序的堆栈、线程、内存等情况

/*
终端交互实时查看
go tool pprof http://localhost:6060/debug/pprof/heap
top 10                 # 查看前10的heap消耗
list 正则表达式         # 查看代码的heap消耗
*/

// 分析文件方式见： test/banchmark
/*
# 生成文件
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

# 查看火焰图、性能等数据
go tool pprof -http=:8080 cpu.prof
*/

func pprofNetHTTP() {
	go func() {
		for {
			time.Sleep(10 * time.Second)
			str := fmt.Sprintf("%v", "hi")
			log.Println(str)
		}
	}()

	http.ListenAndServe(":6060", nil)
}
