
### context
 
文档：[godoc/context](https://godoc.org/context)  

简介：上下文环境, 为每一个进来的请求创造一个context,出去的函数接收这个context, 形成一个请求调用链。

功能： 为请求设置deadline;  传递取消信号; 共享请求的value

实际场景： 
 1. 当子goroutine需要用到父goroutine请求数据时，可用context的value保存传递。
 2. 当父goroutine被取消或超时，可用context的Done通知所有子goroutine，阻止goroutine资源泄露。

注意： 
  - 不要在struct中保存Contexts类型, 函数之间明确传递,最好是第一个参数 `func DoSomething(ctx context.Context, arg Arg) error {	// ... use ctx ...  }
    `
  - 不要传递nil的Context,无法确认是否会使用Context可传递context.TODO
  - context的Vaules仅仅用来保存http请求或者API的数据，不要用来传递函数参数
  - 相同的context可以传递给不同的goroutine，多个goroutine同时发生时context是安全的
  
example [context](https://github.com/golang/blog/tree/master/content/context)：
   
  ```bash
  ### 项目结构与功能
  # 调用google API查询发送的关键字返回结果
  google
    google.go    # goroutine处理搜索,父ctx发送信号后等待goroutine的处理
  server
    server.go    # 程序入口
  userip
    userip.go    # 解析请求IP,设置ctx中value
  ```

  
##### goroutine要处理的问题？

- 数据共享
- 释放回收

##### context  [doc](https://godoc.org/context#Context)   [blog](https://blog.golang.org/context)

###### 解决的问题

why：go服务每个进来的请求都有一个goroutine。请求可能是访问数据库或者RPC服务，当请求取消或超时时，所有的goroutine应该快速回收各自使用的资源。

what：context package处理一个请求的 请求作用域值`value`、取消`Done()`、超时`Deadline()`

Go http的请求，每个请求对应一个goroutine处理，此时该处理函数会启动额外的goroutine处理其他数据，如：数据库和rpc服务。

1. 当子的goroutine访问时需要用到父goroutine数据时，可以用context保存。
2. 当父goroutine被取消或超时时，所有子goroutine都应该退出，释放资源。可以用context的Done()处理。

###### 代码思路梳理 

- [server](https://github.com/golang/blog/blob/master/content/context/server/server.go)

  对于客户端的每个请求创建上下文对象，并设置超时时间`context.WithTimeout(context.Background(), timeout)`

  解析userIP，设置为上下文 `userip.NewContext(ctx, userIP) ` ，google搜索用到上下文 `google.Search(ctx, query) `

  ```go
  var (
  		ctx    context.Context
  		cancel context.CancelFunc
  	)
  timeout, err := time.ParseDuration(req.FormValue("timeout"))
  if err == nil {
  	ctx, cancel = context.WithTimeout(context.Background(), timeout) // 创建context
  } else {
  	ctx, cancel = context.WithCancel(context.Background())
  }
  defer cancel() 
  
  ctx = userip.NewContext(ctx, userIP)    // 使用context传递
  ...
  results, err := google.Search(ctx, query)
  ```

- [userip](https://github.com/golang/blog/blob/master/content/context/userip/userip.go) 

  在context中设置共享值`context.WithValue(ctx, userIPKey, userIP) `，获取值`ctx.Value(userIPKey).(net.IP) `

  ```go
  func NewContext(ctx context.Context, userIP net.IP) context.Context {
  	return context.WithValue(ctx, userIPKey, userIP)      // 设置值
  }
  
  func FromContext(ctx context.Context) (net.IP, bool) {
  	userIP, ok := ctx.Value(userIPKey).(net.IP)         // 获取值
  	return userIP, ok
  }
  ```

- [google](https://github.com/golang/blog/blob/master/content/context/google/google.go)

  并发一个请求完成搜索`go func() { c <- f(client.Do(req)) }() `，保证父退出时，子goroutine退出 `case <-ctx.Done():`

  ```go
  if userIP, ok := userip.FromContext(ctx); ok {
  		q.Set("userip", userIP.String())
  }
  
  func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
  	tr := &http.Transport{}
  	client := &http.Client{Transport: tr}
  	c := make(chan error, 1)
  	go func() { c <- f(client.Do(req)) }()    // 启动goroutine
  	select {
  	case <-ctx.Done():                        // 保证父退出时，子goroutine退出
  		tr.CancelRequest(req)
  		<-c // Wait for f to return.
  		return ctx.Err()
  	case err := <-c:
  		return err
  	}
  }	
  ```

###### [原文：Go Concurrency patterns：context](https://blog.golang.org/context)

###### [翻译：Go语言并发模型：使用 context](https://segmentfault.com/a/1190000006744213)

###### [简书：如何使用context](https://www.jianshu.com/p/0dc7596ba90a)
   
  


  




















  



