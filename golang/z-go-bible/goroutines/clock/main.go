// 并发的clock时钟服务
// 每隔一秒钟将当前时间写到客户端
// 支持启动时指定监听端口
// clock -port 8010
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "set listen port")
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
