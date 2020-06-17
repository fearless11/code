// Package handlers 如何测试内部服务端点的执行效果
package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 运行： go test -v

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	Routes()
}

func TestSendJSON(t *testing.T) {

	t.Log("Given the need to test the SendJSON endpoint")
	{
		req, err := http.NewRequest("GET", "/sendJSON", nil)
		if err != nil {
			t.Fatal("\tshould be able to create a request.", ballotX, err)
		}
		t.Log("\tshould be able to create a request.", checkMark)

		rw := httptest.NewRecorder()
		// DefaultServeMux服务默认的多路由选择器的ServeHTTP方法模拟客户端对/sendJSON的请求
		// 将响应结果对应到recorde中，实现对内容进行解码
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\tshould receive \"200\"", ballotX, rw.Code)
		}

		t.Log("\tshould receive 200", checkMark)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tshould decode the response.", ballotX)
		}
		t.Log("\tshould decode the response.", checkMark)

		if u.Name == "Bill" {
			t.Log("\tshould have a Name.", checkMark)
		} else {
			t.Error("\tshould have a Name.", ballotX, u.Name)
		}

		if u.Email == "www@aa.com" {
			t.Log("\tshould have an Email.", checkMark)
		} else {
			t.Error("\tshould have an Email.", ballotX, u.Email)
		}

	}
}
