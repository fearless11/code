

### context
 
[godoc/context](https://godoc.org/context)  

#### 简介
  
   上下文环境, 为每一个进来的请求创造一个context,出去的函数接收这个context, 形成一个请求调用链。

#### 功能
 为请求设置deadline;  传递取消信号; 共享请求的value

 实际场景： 

 1. 当子goroutine需要用到父goroutine请求数据时，可用context的value保存传递。
 2. 当父goroutine被取消或超时，可用context的Done通知所有子goroutine，阻止goroutine资源泄露。

#### 规则
  - 不要在struct中保存Contexts类型, 函数之间明确传递,最好是第一个参数
    ```go
    func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
    }
    ```
  - 不要传递nil的Context,无法确认是否会使用Context可传递context.TODO
  - context的Vaules仅仅用来保存http请求或者API的数据，不要用来传递函数参数
  - 相同的context可以传递给不同的goroutine，多个goroutine同时发生时context是安全的
  


#### example
   
- 项目结构  [blog/context](https://github.com/golang/blog/tree/master/content/context)

  ```bash
  # 实现调用google API查询发送的关键字返回结果
  google
    google.go    # goroutine处理搜索,父ctx发送信号后等待goroutine的处理
  server
    server.go    # 程序入口
  userip
    userip.go    # 解析请求IP,设置ctx中value


  # 运行项目后发送POST的Form请求, q:搜索内容(必选); timeout:10 (可选)
  ```
   
  


  




















  



