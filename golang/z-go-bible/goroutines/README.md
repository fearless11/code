#### 使用
`go f() // create a new goroutine that calls f(); don't wait`

#### 控制goroutine方法
- 计数信号量
  ```
  // sema is a counting semaphore for limiting concurrency 
  var sema = make(chan struct{}{},20)

  func xxx(){
      sema <- struct{}{}           // acquire token
      defer func(){  <- sema }()   // release token
  }
  ```

#### 为了等待所有goroutine结束
- sync
   ```
   var wg sync.WaitGroup      // number of working goroutines
   for f := range filenames {
        wg.Add(1)
        // worker
        go func() {
            defer wg.Done()
            ......
        }()
   }

    // closer
    go func() {
        wg.Wait()
        .....
    }()
   ```
#### goroutine如何并发的退出

思路: 广播机制。不要向channel发送值，而是用关闭一个channel来进行广播。
```
var done = make(chan struct{})

func cancelled() bool {
    select {
        case <-done:
            return true
        default:
            return false
    }
}

go func() {
    os.Stdin.Read(make([]byte, 1)) // read a single byte
    close(done)
}()
```

#### goroutine和线程的本质区别