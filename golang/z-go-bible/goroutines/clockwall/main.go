// 连接clock服务的客户端
// 类似 nc ip port
// 支持一次显示所有服务器传回的结果
// clockwall NewYork=localhost:8010 Tokyo=localhost:8020

package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	for _, arg := range os.Args[1:] {
		clock := strings.Split(arg, "=")
		go netcat(clock[1])
	}
	select {}
}

func netcat(server string) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
