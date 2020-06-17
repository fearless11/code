// 使用通道来监视程序运行的时间
// 以在程序运行时间过长时如何终止程序
package main

import (
	"log"
	"os"
	"time"

	"code/goroutine/mode/runner/runner"
)

//////// 控制程序生命周期，goroutine超时退出

const timeout = 3 * time.Second

func main() {
	log.Println("starting work.")

	r := runner.New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("process  ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("process - task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
