package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ArithRequest struct {
	A int
	B int
}

type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func netrpc() {
	// 基于http的rpc
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}

	req := ArithRequest{1, 2}
	var res ArithResponse

	err = conn.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("netrpc %d * %d = %d\n", req.A, req.B, res.Pro)
}

func netrpcJSON() {
	// 基于tcp的jsonrpc
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalln("conn", err)
	}

	req := ArithRequest{1, 2}
	var res ArithResponse

	err = conn.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatalln("call", err)
	}

	fmt.Printf("netrpcJSON %d * %d = %d\n", req.A, req.B, res.Pro)
}

func main() {
	netrpc()
	netrpcJSON()
}
