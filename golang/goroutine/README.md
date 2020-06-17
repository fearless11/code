<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [什么是并发？](#%E4%BB%80%E4%B9%88%E6%98%AF%E5%B9%B6%E5%8F%91)
  - [并发&并行](#%E5%B9%B6%E5%8F%91%E5%B9%B6%E8%A1%8C)
  - [进程&线程&goroutine](#%E8%BF%9B%E7%A8%8B%E7%BA%BF%E7%A8%8Bgoroutine)
  - [进程&线程](#%E8%BF%9B%E7%A8%8B%E7%BA%BF%E7%A8%8B)
  - [线程&goroutine](#%E7%BA%BF%E7%A8%8Bgoroutine)
  - [进程间通信方式](#%E8%BF%9B%E7%A8%8B%E9%97%B4%E9%80%9A%E4%BF%A1%E6%96%B9%E5%BC%8F)
  - [goroutine配置](#goroutine%E9%85%8D%E7%BD%AE)
  - [goroutine的GMP模型](#goroutine%E7%9A%84gmp%E6%A8%A1%E5%9E%8B)
  - [线程调度模型](#%E7%BA%BF%E7%A8%8B%E8%B0%83%E5%BA%A6%E6%A8%A1%E5%9E%8B)
- [并发要解决的问题？](#%E5%B9%B6%E5%8F%91%E8%A6%81%E8%A7%A3%E5%86%B3%E7%9A%84%E9%97%AE%E9%A2%98)
  - [竞争状态](#%E7%AB%9E%E4%BA%89%E7%8A%B6%E6%80%81)
  - [解决竞争办法](#%E8%A7%A3%E5%86%B3%E7%AB%9E%E4%BA%89%E5%8A%9E%E6%B3%95)
  - [并发模式](#%E5%B9%B6%E5%8F%91%E6%A8%A1%E5%BC%8F)
  - [通道](#%E9%80%9A%E9%81%93)
    - [无缓冲通道](#%E6%97%A0%E7%BC%93%E5%86%B2%E9%80%9A%E9%81%93)
    - [有缓冲通道](#%E6%9C%89%E7%BC%93%E5%86%B2%E9%80%9A%E9%81%93)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


## 什么是并发？

### 并发&并行

- 并发: 同时管理很多事情，有些事可能做一半停止做其他；一次买两个包子

- 并行: 同时做很多事情，不同代码片在同时在不同处理器上执行；一口吃两个包子

### 进程&线程&goroutine

- 进程：操作系统分配资源的基本单位，资源有内存空间、文件和设备句柄以及线程等

- 线程：操作系统运算和调度的最小单位，调度处理程序中的代码

- goroutine：go逻辑处理器调度的基本单位，调度处理go代码
  
### 进程&线程

- 对应关系：一个线程只能属于一个进程，一个进程可以有多个线程
- 系统资源：系统资源分配给进程，处理器分配给线程
- 并发性：不同进程间可以并发，线程也可以
- 系统开销：创建、撤销进程比线程的开销大
  
### 线程&goroutine
  
- 对应关系: 一个线程绑定一个逻辑处理器，一个逻辑处理器管理多个goroutine
- 灵活性:   线程由内核管理与调度，goroutine由go调度器调度处理。不通过硬件时钟来定期触发调度
- 系统开销: goroutine的调度开销远远小于线程调度开销
- 栈空间:   OS的线程都有一个固定大小的栈内存(2MB)，goroutine在栈内存是按需增大和缩小，最大限制可到1GB

  [线程和gorutine区别](http://www.cnblogs.com/yanghuahui/p/9043631.html)

### 进程间通信方式

- 管道、信号、消息队列、共享内存、信号量、套接字

### goroutine配置
  
```go
// 修改任何语言运行时配置参数的时候，需要配合基准测试来评估程序运行结果

// 设置处理器核数
runtime.GOMAXPROCS(1)
```

### goroutine的GMP模型

- G：goroutine，M：线程， P：处理器，


### 线程调度模型
- 进程，操作系统，线程

## 并发要解决的问题？

### 竞争状态
没有同步情况下，对共享资源进行读写的状态
```bash
 # 竞争检测器标志来编译程序，后运行程序，查看具体竞争信息
 go build -race
```
### 解决竞争办法

- 原子函数（整型）

- 互斥锁（临界区）

- 通道（无竞争状态，通信共享）

### 并发模式

- runner : 控制程序生命周期，goroutine超时退出
- pool : 管理可复用的资源池，用有缓冲的通道记录资源的多少
- work : 管理goroutine数量，类似消费者-生产者模式，动态创建goroutine消费者处理任务

### 通道

不要通过共享内存来通信，而应该通过通信来共享内存

```go
// 无缓冲的整型通道
unbuffered := make(chan int)
// 有缓冲的字符串通道
buffered := make(chan string, 10)

// 向通道发送一个字符串
buffered <- "Gopher"
// 从通道接收一个字符串
value := <-buffered
```

#### 无缓冲通道

通过发送和接收需要共享的资源，在 goroutine 之间做同步，保证同时交换数据

#### 有缓冲通道 

关闭通道后只能从通道接收数据 不能向通道发送数据

 ```go
  resource := make(chan io.reader,4)
  
  // 清空通道前，先关闭否则会死锁
  close(resource)

  // 没有数据for阻塞，通道关闭才退出
  for r := range resource {    
      r.close()
  }

// 判断通道状态
task, ok := <-tasks
if !ok {
  // 说明通道清空且已关闭
	fmt.Println("shutting down")
	return
}
```