// Package handlers 提供用于网络服务的服务端点
// 包名可以和目录名不一样
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Routes 为网络服务设置路由
func Routes() {
	http.HandleFunc("/sendJSON", SendJSON)
}

// SendJSON 返回一个JSON文档
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "www@aa.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}

func noPublicFunc() {
	fmt.Println("test example for no public function")
}
