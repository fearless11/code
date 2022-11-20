<!-- TOC -->

- [pprof](#pprof)
  - [简介](#简介)
  - [实现方式](#实现方式)
  - [采集数据](#采集数据)
    - [基本测试](#基本测试)
    - [runtime/pprof](#runtimepprof)
    - [net/http/pprof](#nethttppprof)
  - [分析数据](#分析数据)
    - [终端交互式](#终端交互式)
    - [web界面](#web界面)
    - [工具处理](#工具处理)

<!-- /TOC -->



## pprof

[google-pprof](https://github.com/google/pprof)

[runtime/pprof](https://godoc.org/runtime/pprof)

[net/http/pprof](https://godoc.org/net/http/pprof)

[Golang特性剖析pprof](https://www.jianshu.com/p/4e4ff6be6af9)

### 简介

    一个可视化、分析程序运行时性能数据的工具

### 实现方式

- 采集文件后通过`go tool pprof` 或 google的`pprof`工具分析
- 通过http接口进入终端实时分析

### 采集数据

#### 基本测试

```bash
 # 产生cpu.prof 和 mem.prof文件
 go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
```

#### runtime/pprof
```go
// 产生cpuprofile和memprofile文件
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }

    // ... rest of the program ...

    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
    }
}
```
#### net/http/pprof

```go
// 提供接口 http://localhost:6060/debug/pprof/ 
import _ "net/http/pprof"

go func() {
	log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

### 分析数据

#### 终端交互式

```bash
# look at the heap profile
go tool pprof http://localhost:6060/debug/pprof/heap

# look at a 30-second CPU profile:
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

# look at the goroutine blocking profile, after calling runtime.SetBlockProfileRate in your program:
go tool pprof http://localhost:6060/debug/pprof/block

# collect a 5-second execution trace
wget http://localhost:6060/debug/pprof/trace?seconds=5

# look at the holders of contended mutexes, after calling runtime.SetMutexProfileFraction in your program:
go tool pprof http://localhost:6060/debug/pprof/mutex
```

#### web界面
  
  安装 `Graphviz`  [baiduyun]( https://pan.baidu.com/s/1czQxYY?fid=393219168683528 )
  
  ```bash
  官网注册才能下载，可用百度云

 配置PATH： 电脑 —> 属性 —> 高级属性设置 —>环境变量 —> 编辑Path加入Graphviz目录\bin

 win+R后cmd验证 dot -version
  ```

- 接口方式

```bash
# web browser
http://localhost:6060/debug/pprof/

# pprof can read a profile from a file or directly from a server via http.
# 可下载profile.tgz
http://localhost:6060/debug/pprof/profile
```

- 分析文件
```bash
 # 指定文件后,可多种方式查看top、调用关系图、火焰图等
 go tool pprof -http=:8080 cpu.prof
```

#### 工具处理
```bash
# 安装工具
go get -u github.com/google/pprof

# Generate a text report of the profile, sorted by hotness
pprof -top [main_binary] profile.pb.gz

# Generate a graph in an SVG file, and open it with a web browser
pprof -web [main_binary] profile.pb.gz

# Run pprof on interactive mode
pprof [main_binary] profile.pb.gz

# Run pprof via a web interface
pprof -http=[host]:[port] [main_binary] profile.pb.gz
```

