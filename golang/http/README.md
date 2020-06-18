<!-- TOC -->

- [HTTP协议](#http协议)
- [net/http包](#nethttp包)
    - [客户端](#客户端)
    - [服务端](#服务端)

<!-- /TOC -->

## HTTP协议

- 传输层协议：底层TCP
- C/S模型： 客户端-服务端
- 特点： 简单快速、无连接、无状态
- 格式： 请求行、请求头部(KV)、空行、请求数据
- 请求方式：[RFC 7231#section-4.3](https://tools.ietf.org/html/rfc7231#section-4.3)
  
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


## net/http包

### 客户端

- 发出请求
 
  `Get, Head, Post, and PostForm`
  ```go
  resp, err := http.Get("http://example.com/")

  resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)

  resp, err := http.PostForm("http://example.com/form",url.Values{"key": {"Value"}, "id": {"123"}})
  ```

- 请求头设置
  
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
  
  ```go
  // clients和Transports是并发安全的、能被重复使用
  // request.Body必须关闭
  resp, err := http.Get("http://example.com/")
  if err != nil {
    // handle error
  }
  // 关闭body
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  ```

### 服务端

- 监听与处理
  
  ```go
  // handle添加到DefaultServeMux
  http.Handle("/foo", fooHandler)
  // HandleFunc添加到DefaultServeMux
  http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
  ```

- 设置超时
  
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