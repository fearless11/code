
<!-- TOC -->

- [RPC](#rpc)
    - [简介](#简介)
    - [go支持的包](#go支持的包)
    - [具体过程](#具体过程)

<!-- /TOC -->

## RPC

- [Go语言RPC协议](http://c.biancheng.net/view/4522.html)

### 简介

- 远程过程调用（Remote Procedure Call，简称 RPC）是一个计算机通信协议
- RPC 协议基于TCP、UDP 或 HTTP之上
- RPC 允许跨机器、跨语言调用计算机程序
- RPC 允许运行于一台计算机的程序调用另一台计算机的子程序方法或函数


### go支持的包

-  net/rpc 包： 基于tcp或http， 采用encoding/gob 编解码，调用可跨机器，一般不跨语言。
-  net/rpc/jsonrpc 包： 基于tpc， 采用 JSON 编解码，调用可跨机器，跨语言。
-  第三方 rpc 包： 采用 protobuf 编解码，根据 protobuf 声明文件自动生成 rpc 方法定义与服务注册代码，方便的进行 rpc 服务调用。

### 具体过程

- 服务端： 对象注册，允许对象的导出方法就可以被远程访问。
  ```bash
  # 导出方法的标准
  1. 方法有两个参数，都是导出类型或内建类型；
  2. 方法的第二个参数是指针类型；
  3. 方法只有一个 error 接口类型的返回值。
  ```

- 客户端： 调用call对象的方法执行



  




