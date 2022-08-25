[toc]

### 简介
- 2007年开始，2009年11月官方发布。
- Google开源，由Robert Griesemer，Rob Pike 和Ken Thompson开发。

### 优点

- 上手快，入门简单，学习成本低
- 原生并发

### 特点

- 静态型编译，打包二进制后即可运行
- goroutine协程，一个线程可执行多个goroutine，由go的逻辑处理器调度，资源占用少可高并发
- channel通道，不要通过共享内存来通信，而应该通过通信来共享内存
- 函数是一等公民
- 组合结构体，实现代码复用，不同结构体之间的合作
- 接口，实现接口的所有方法则实现了接口本身

### 风格

- 变量名：首字符大写
- 常量名：大写字母
- 包名：小写单词，不用下划线或驼峰记法
- 文件名：小写，可加下划线分割
- 结构体名：驼峰法
- 接口名：只包含一个方法的接口的名称加上-er后缀Reader
- 驼峰记法：驼峰记法 MixedCaps 或 mixedCaps
- 左括号：不应将一个控制结构（if、for、switch 或 select）的左大括号放在下一行

### 资料

- 网站

[Go指南 tour.go-zh.org](https://tour.go-zh.org/list)

[Go例子 gobyexample.com](https://gobyexample.com/)

[Go实效编程-effective_go](https://go-zh.org/doc/effective_go.html)

[Go官网 pkg.go.dev](https://pkg.go.dev/)

[Go文档中文 go-zh.org](https://go-zh.org/doc)

[CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)

[Go代码风格](https://github.com/golang-standards/project-layout)

[go语言中文网 studygolang.com](https://studygolang.com/dl)

[代理-goproxy.io](https://goproxy.io/)

[命令-go-zh.org/cmd/go](https://go-zh.org/cmd/go/)

- 书籍

[Go实效编程](https://bingohuang.gitbooks.io/effective-go-zh-en/)

[Go语言圣经](https://docs.hacknode.org/gopl-zh/ch0/ch0-01.html)

[the-way-to-go中文](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)

- 博客

[如何写出优雅的golang代码](https://draveness.me/golang-101/)

[Go文档汇集](https://www.topgoer.com/)

[飞雪无情 www.flysnow.org](https://www.flysnow.org/)

[狼人 blog.wolfogre.com](https://blog.wolfogre.com)

[力扣-github.com/aQuaYi/LeetCode-in-Go](https://github.com/aQuaYi/LeetCode-in-Go)

[Go核心36讲-极客时间-郝林](https://account.geekbang.org/dashboard/buy)

- 框架&实用

[网站-gin](https://github.com/gin-gonic/gin)

[网站-beego](https://github.com/astaxie/beego)

[数据库-gorm](https://gorm.io/docs/query.html) 

[文档管理-gin-swagger](https://github.com/swaggo/gin-swagger)

[文档管理-redoc](https://github.com/Redocly/redoc)

[https://mholt.github.io/json-to-go/](https://mholt.github.io/json-to-go/)

Mock单元测试 [gomock](https://github.com/golang/mock) [sqlmock](https://github.com/DATA-DOG/go-sqlmock)  [httpmock](https://github.com/jarcoal/httpmock) [monkey](https://github.com/bouk/monkey)

- 项目

[beats社区](https://www.elastic.co/guide/en/beats/devguide/7.0/index.html)

[grafana](https://github.com/grafana/grafana/blob/master/pkg/login/ldap_login.go)

### 安装
- 下载golang安装包 [https://gomirrors.org/](https://gomirrors.org/)

  - windows
  
     ```bash
     # 下载 
     https://studygolang.com/dl/golang/go1.14.4.windows-amd64.zip
     # 解压路径
     D:\Program Files (x86)\
     # 添加windows变量
     桌面 -> 我的电脑(右击) -> 属性 -> 高级系统设置 -> 环境变量 -> 用户变量 -> 新建：
       GOROOT D:\Program Files (x86)\go
       GOPATH D:\gohome
       GO111MODULE on
       GOPROXY https://mirrors.aliyun.com/goproxy/
     # 编辑Path变量
     Path变量编辑 -> 新建：
     %GOROOT%\bin
     %GOPATH%\bin
     # 命令查看验证
     win+R -> cmd -> go version
     ```
  - mac
    ```bash
    # 下载
    wget https://studygolang.com/dl/golang/go1.14.4.darwin-amd64.tar.gz
    tar xf go1.14.4.darwin-amd64.tar.gz -C /usr/local
    mkdir -p /data/go/src
    # 配置
    # root用户
    cp /etc/profile /home
    echo 'export GOROOT=/usr/local/go' >> /etc/profile
    echo 'export GOPATH=/data/go' >> /etc/profile
    echo 'export GO111MODULE=on' >> /etc/profile
    echo 'export GOPROXY=https://mirrors.aliyun.com/goproxy/' >> /etc/profile
    echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> /etc/profile
    source /etc/profile
    # 普通用户
    cp ~/.bash_profile /home
    echo 'export GOROOT=/usr/local/go' >> ~/.bash_profile
    echo 'export GOPATH=/data/go' >> ~/.bash_profile
    echo 'export GO111MODULE=on' >> ~/.bash_profile
    echo 'export GOPROXY=https://mirrors.aliyun.com/goproxy/' >> ~/.bash_profile
    echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> ~/.bash_profile
    echo 'Source ~/.bash_profile' >  ~/.zshrc
    source ~/.zshrc
    # 验证
    go env
    ```

  - linux
    
    ```bash
    wget https://studygolang.com/dl/golang/go1.14.4.linux-amd64.tar.gz
    tar xf go1.14.4.linux-amd64.tar.gz -C /usr/local
    mkdir -p /opt/go/src
    cp /etc/profile /home
    echo 'export GOROOT=/usr/local/go' >> /etc/profile
    echo 'export GOPATH=/opt/go' >> /etc/profile
    echo 'export GOPROXY=https://mirrors.aliyun.com/goproxy/' >> /etc/profile
    echo 'export PATH=$PATH:$GOROOT/bin:$GOPATH/bin' >> /etc/profile
    source /etc/profile
    go env
    ```
- 下载vscode编辑器 [https://code.visualstudio.com/Download](https://code.visualstudio.com/Download)

  - 安装golang插件并配置

    ```bash
    # 安装
    快捷键 F1或ctrl+shift+p ——> 输入命令 Extensions:Install Extension ——> 插件管理搜索go后安装

    # 配置
    菜单 — Preferences — User - Extensions - Go - Edit in settings.json
    # windows环境 用户空间  user/settings.json
    {
        "go.goroot": "/usr/local/go",
        "go.gopath": "/data/go",
        "go.buildOnSave": "package",
        "go.lintOnSave": "package",
        "go.formatTool": "goimports",      
        "go.gocodeAutoBuild": false,
        "go.useGoProxyToCheckForToolUpdates":true 
    }
    # 工作空间 F5进行调试，调试配置 .vscode/launch.json
    {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Launch",
                "type": "go",
                "request": "launch",
                "mode": "auto",
                "stopOnEntry": false,
                "program": "${fileDirname}",
                "env": {},
                "args": []
            }
        ]
    }
    ```

  - 安装go的工具 [https://github.com/golang](https://github.com/golang)

    ```bash
    # vscode会提示下载安装，会出现安装失败的
    
    # 直接手动下载安装
    go get -u -v  github.com/mdempsky/gocode 
    go get -u -v  github.com/uudashr/gopkgs/cmd/gopkgs 
    go get -u -v  github.com/ramya-rao-a/go-outline  
    go get -u -v  github.com/acroca/go-symbols  
    go get -u -v  github.com/go-delve/delve/cmd/dlv  
    go get -u -v  github.com/rogpeppe/godef
    go get -u -v  github.com/sqs/goreturns  
    go get -u -v  github.com/cweill/gotests
    go get -u -v  github.com/godoctor/godoctor
    
    # 被墙的，先在下载再安装。也可以设置代理解决
    mkdir -p $GOPATH/src/golang.org/x
    cd $GOPATH/src/golang.org/x
    git clone https://github.com/golang/tools.git tools
    git clone https://github.com/golang/lint.git lint
    git clone https://github.com/golang/net.git net
    
    go install golang.org/x/tools/cmd/guru
    go install golang.org/x/tools/cmd/gorename
    go install golang.org/x/lint/golint
    go install golang.org/x/net/gonet
    ```
