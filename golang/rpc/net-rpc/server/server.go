package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 使用http作为RPC的载体

// 只有满足如下标准的方法才能用于远程访问，其余方法会被忽略：
// 方法是可导出的；
// 方法有两个参数，都是导出类型或内建类型；
// 方法的第二个参数是指针类型；
// 方法只有一个 error 接口类型的返回值。

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

type Arith struct {
}

func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除以零")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func netRPC() {
	// 对象注册
	rpc.Register(new(Arith))

	// 基于http
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("start listen netRPC 127.0.0.1:8080 ")

	http.Serve(lis, nil)
}

func netJSONRPC() {

	// 对象注册
	rpc.Register(new(Arith))

	lis, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("start listen netJSONRPC 127.0.0.1:8081")

	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}

		go func(conn net.Conn) {
			fmt.Println("netJSONRPC receiver connect")
			// 基于tcp
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}

func main() {

	// 基于http或tcp协议, 采用Go特有的 gob 编码协议, 一般不支持跨语言调用。
	go netRPC()

	// 基于tcp协议, 采用 JSON 进行数据编解码, 支持跨语言调用。
	netJSONRPC()
}
