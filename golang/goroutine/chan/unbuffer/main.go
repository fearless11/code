// 无缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何值的通道
//  发送和接收同时准备好，否则其中一个会阻塞等待

//////////////  2个goroutine间的网球比赛 /////////////
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// go run main.go
/* output
palyer Djokovic hit 1
palyer Nadal hit 2
palyer Djokovic hit 3
palyer Nadal hit 4
Player Djokovic missed
player Nadal won
*/

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	wg.Add(2)

	// 启动两个球手
	go player("Nadal", court)
	go player("Djokovic", court)

	// 发球
	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 等待接球
		ball, ok := <-court
		if !ok {
			// 如果通道关闭,我们就赢了
			fmt.Printf("player %s won\n", name)
			return
		}

		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s missed\n", name)
			// 关闭通道，表示我们输了
			close(court)
			return
		}

		// 显示击球数，并将击球数加 1
		fmt.Printf("palyer %s hit %d\n", name, ball)
		ball++

		// 将球打向对手
		court <- ball
	}
}
