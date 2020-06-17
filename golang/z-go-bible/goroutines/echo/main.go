// 并发的Echo服务
// 客户端发送的内容返回,像回声一样
// 客户端发送：Hello
// 服务段返回：HELLO
// 			 Hello
// 			 hello

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var port string
var wg sync.WaitGroup

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
		wg.Add(1)
		go handleConn(conn)

		go func() {
			wg.Wait()
			log.Println("a grocer typist exit")
		}()
	}

}

func handleConn(c net.Conn) {
	defer wg.Done()
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))

}
