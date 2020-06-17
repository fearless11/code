[toc]


#### 简介

提供服务端、客户端

#### 使用

##### 客户端
- 发出请求
 
  `Get, Head, Post, and PostForm`
  ```go
  resp, err := http.Get("http://example.com/")

  resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)

  resp, err := http.PostForm("http://example.com/form",url.Values{"key": {"Value"}, "id": {"123"}})
  ```

- 请求设置
  
  `request` 设置头部、重定向策略或其他
  ```go
  // client对象
  client := &http.Client{
	  CheckRedirect: redirectPolicyFunc,
  }
  // 方法
  resp, err := client.Get("http://example.com")
  // 请求头部设定
  req, err := http.NewRequest("GET", "http://example.com", nil)
  req.Header.Add("If-None-Match", `W/"wyzzy"`)
  // 发送
  resp, err := client.Do(req)
  ```

- 传输设置

  `Transport` 设置proxies、TLS configuration、keep-alives、compression等
  ```go
  tr := &http.Transport{
	  MaxIdleConns:       10,
	  IdleConnTimeout:    30 * time.Second,
	  DisableCompression: true,
  }

  client := &http.Client{Transport: tr}
  resp, err := client.Get("https://example.com")
  ```

- 注意
  
  `clients`和`Transports`是并发安全的、能被重复使用
  `request.Body`必须关闭
  ```go
  resp, err := http.Get("http://example.com/")
  if err != nil {
    // handle error
  }
  // 关闭body
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  ```

##### 服务端

- 监听与处理
  
  a given address and handler. 
  The handler is usually nil, which means to use DefaultServeMux. 
  ```go
  // handle添加到DefaultServeMux
  http.Handle("/foo", fooHandler)
  // HandleFunc添加到DefaultServeMux
  http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
  ```
- 设置
  
  ```go
  s := &http.Server{
	  Addr:           ":8080",
	  Handler:        myHandler,
	  ReadTimeout:    10 * time.Second,
	  WriteTimeout:   10 * time.Second,
  	MaxHeaderBytes: 1 << 20,
  }
  log.Fatal(s.ListenAndServe())
  ```
#### 源码

##### 常量 
  [RFC 7231#section-4.3](https://tools.ietf.org/html/rfc7231#section-4.3)
  ```
  const (
    MethodGet     = "GET"
    MethodHead    = "HEAD"
    MethodPost    = "POST"
    MethodPut     = "PUT"
    MethodPatch   = "PATCH" // RFC 5789
    MethodDelete  = "DELETE"
    MethodConnect = "CONNECT"
    MethodOptions = "OPTIONS"
    MethodTrace   = "TRACE"
  )
  ```

##### ServeMux


- 



### context


