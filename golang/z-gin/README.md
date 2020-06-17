### Gin

[github-Gin](https://github.com/gin-gonic/gin)

- 定义：一个web框架，性能比 httprouter快40倍
- 本质：在net/http上针对handler层做操作
- 具体：
	- 路由: 精确匹配、模糊匹配
  - 路由分组
	- 中间件（上下文日志、认证、防止崩溃）
  - 绑定：对各种格式的参数进行序列化,响应的反序列化，获取cookie
  - 模板渲染：响应进行模板渲染
  - 多个http监听
  - 第三方实现优雅退出：当收到系统信号时，先等待context上下文都退出才结束程序

- 使用

  ```bash
  cd gin-web
  go mod init gin-web
  go get -x -u github.com/gin-gonic/gin
  
  import "github.com/gin-gonic/gin"
  ```