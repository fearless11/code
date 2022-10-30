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
