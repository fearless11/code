// 示例测试
// 既能用于测试，也能用于文档
// 示例函数必须是公开的，首字母大写
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

// 运行： go test -v

// 初始化运行服务器，否则ExampleSendJSON访问为404
func init() {
	Routes()
}

// ExampleSendJSON 可以显示在Go文档里
//
func ExampleSendJSON() {
	r, _ := http.NewRequest("GET", "/sendJSON", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, r)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println("ERROR", err)
	}

	fmt.Println(u)
	//  Output:
	// {Bill www@aa.com}
}

// 运行：go test -v -run="ExamplenoPublicFunc"
// output: testing: warning: no tests to run
func ExamplenoPublicFunc() {
	noPublicFunc()
}
