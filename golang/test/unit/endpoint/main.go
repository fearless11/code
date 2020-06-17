// 实现简单的网络服务
// 调用服务端点
// 在构造网络API，希望直接测试自己的服务的所有服务端点，而不是启动整个网络服务
// 类似mock机制
package main

import (
	"log"
	"net/http"

	"gitee.com/feareless11/golang/code/test/unit/endpoint/handlers"
)

// 运行： go run main.go
// 浏览器访问： http://localhost:4000/sendJSON
// output:
// {"Name":"Bill","Email":"www@aa.com"}

func main() {
	handlers.Routes()

	log.Println("listener : start : Listening on: 4000")
	http.ListenAndServe(":4000", nil)
}
