### channel特点
- 用于通信
- 有特殊类型
- 引用类型，零值nil
- 不带缓冲channel用于实现同步
- 关闭的channel发送数据会panic
- 只有发送者关闭channel
- 带缓冲的channel等价于元素队列

### channal操作
```
口诀： <-chan chan里读取
      chan<- 发送chan里
ch := make(chan int,0)   //定义
ch <- x                  //接收
x = <-ch                 //发送
<-ch                     //发送
close(ch)            //关闭通道
```

### 内存泄露
```
func mirroredQuery() string {
    responses := make(chan string, 3)
    go func() { responses <- request("asia.gopl.io") }()
    go func() { responses <- request("europe.gopl.io") }()
    go func() { responses <- request("americas.gopl.io") }()
    return <-responses // return the quickest response
}

func request(hostname string) (response string) { /* ... */ }
```
分析
```
如果我们使用了无缓存的channel，那么两个慢的goroutines将会因为没有人接收而被永远卡住。这种情况，称为goroutines泄漏，这将是一个BUG。和垃圾变量不同，泄漏的goroutines并不会被自动回收，因此确保每个不再需要的goroutine能正常退出是重要的。
```
缓冲长度
```
关于无缓存或带缓存channels之间的选择，或者是带缓存channels的容量大小的选择，都可能影响程序的正确性。无缓存channel更强地保证了每个发送操作与相应的同步接收操作；但是对于带缓存channel，这些操作是解耦的。同样，即使我们知道将要发送到一个channel的信息的数量上限，创建一个对应容量大小的带缓存channel也是不现实的，因为这要求在执行任何接收操作之前缓存所有已经发送的值。如果未能分配足够的缓存将导致程序死锁。

Channel的缓存也可能影响程序的性能。
```