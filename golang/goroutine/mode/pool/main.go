// 使用pool/pool.go来共享一组模拟的数据库连接
package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"code/goroutine/mode/pool/pool"
)

// 代码中资源池的大小并没有限制,只要申请,资源次没有了就直接新创建了。。。。

// go run main.go
/*
2020/03/13 14:43:37 Acquire: new resource
2020/03/13 14:43:37 create: new connection 1
2020/03/13 14:43:37 Acquire: new resource
2020/03/13 14:43:37 Acquire: new resource
2020/03/13 14:43:37 Acquire: new resource
2020/03/13 14:43:37 Acquire: new resource
2020/03/13 14:43:37 create: new connection 2
2020/03/13 14:43:37 create: new connection 3
2020/03/13 14:43:37 create: new connection 4
2020/03/13 14:43:37 create: new connection 5
2020/03/13 14:43:37 QID[2] CID[4]
2020/03/13 14:43:37 release: in Queue
2020/03/13 14:43:37 QID[4] CID[1]
2020/03/13 14:43:37 release: in Queue
2020/03/13 14:43:37 QID[0] CID[5]
2020/03/13 14:43:37 release: closing
2020/03/13 14:43:37 close: connection 5
2020/03/13 14:43:38 QID[3] CID[3]
2020/03/13 14:43:38 release: closing
2020/03/13 14:43:38 close: connection 3
2020/03/13 14:43:38 QID[1] CID[2]
2020/03/13 14:43:38 release: closing
2020/03/13 14:43:38 close: connection 2
2020/03/13 14:43:38 shutdown program
2020/03/13 14:43:38 close: connection 4
2020/03/13 14:43:38 close: connection 1
*/

const (
	//5个goroutine
	maxGoroutines = 5
	// 资源池大小为2
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

// Close 实现了 io.Closer 接口，以便 dbConnection可以被池管理
// Close 用来完成任意资源的释放管理
func (dbConn *dbConnection) Close() error {
	log.Println("close: connection", dbConn.ID)
	return nil
}

var idCounter int32

// createConnection 是一个工厂函数
// 当需要一个新连接时，资源池会调用这个函数
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("create: new connection", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("shutdown program")
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	// 接口实例化: conn.(*dbConnection)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
