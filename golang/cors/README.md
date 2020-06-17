
<!-- TOC -->

- [CORS (Corss-origin resource sharing) 跨域共享资源](#cors-corss-origin-resource-sharing-跨域共享资源)
    - [不限制CORS的危害？](#不限制cors的危害)
    - [限制CORS情况？](#限制cors情况)
    - [谁限制了跨域？](#谁限制了跨域)
    - [跨域解决办法](#跨域解决办法)
- [OPTIONS请求](#options请求)
    - [options出现的情况](#options出现的情况)
    - [复杂的跨域请求浏览器才发送`preflight request`](#复杂的跨域请求浏览器才发送preflight-request)
    - [为什么跨域的复杂请求需要`preflight request`？](#为什么跨域的复杂请求需要preflight-request)

<!-- /TOC -->

### CORS (Corss-origin resource sharing) 跨域共享资源

W3C标准,是一种机制。

同源条件：两个URL的协议、域名和端口必须都相同

浏览器采用同源策略为安全性考虑实施的安全策略。

从一个域上加载的脚本不允许访问另外一个域的文档属性。

#### 不限制CORS的危害？

举个例子：比如一个恶意网站的页面通过iframe嵌入了银行的登录页面（二者不同源）, 如果没有同源限制，恶意网页上的javascript脚本就可以在用户登录银行的时候获取用户名和密码。

#### 限制CORS情况？

浏览器中，`<script>、<img>、<iframe>、<link>`等标签都可以加载跨域资源，而不受同源限制。

浏览器会限制脚本中发起的跨域请求。比如，使用`XMLHttpRequest`对象和`Fetch`发起`HTTP`请求就必须遵守同源策略。

Web应用程序通过`XMLHttpRequest`对象或`Fetch`能且只能向同域名的资源发起`HTTP`请求，而不能向任何其它域名发起请求。

#### 谁限制了跨域？

- 浏览器直接拦截请求

    有些浏览器不允许从`HTTPS`跨域访问`HTTP`，比如`Chrome`和`Firefox`，这些浏览器在请求还未发出拦截

- 浏览器没有从后端收到允许跨域的响应拦截请求

    不允许跨域访问并非是浏览器限制了发起跨站请求，而是跨站请求可以正常发起，但是返回结果被浏览器拦截

    最好的例子是CSRF跨站攻击原理，请求是发送到了后端服务器,无论是否设置允许跨域。

#### 跨域解决办法

- [jsonp协议](https://www.cnblogs.com/dowinning/archive/2012/04/19/json-jsonp-jquery.html)

  原因：Web页面上调用js文件时则不受是否跨域的影响，拥有"src"这个属性的标签都拥有跨域的能力；

  实现：浏览器客户端在带有src上向不同的站点发送GET请求，同时传送callback将返回值作为返回函数的参数

- WebSocket

  全双工，允许服务端push数据给客户端。
- Nginx代理设置

  ```bash
  location / {
   if ($request_method = 'OPTIONS') {
       add_header Access-Control-Allow-Origin $http_origin always;
       add_header Access-Control-Allow-Credentials true always;
       add_header Access-Control-Allow-Methods 'GET,POST,PUT,DELETE,OPTIONS' always;
       add_header Access-Control-Allow-Headers 'Authorization,X-Requested-With,Content-Type,Origin,Accept' always;
       add_header Access-Control-Max-Age 3600;
       add_header Content-Length 0;
       return 200;
   }

   add_header Access-Control-Allow-Origin $http_origin always;
   add_header Access-Control-Allow-Credentials true always;
   add_header Access-Control-Allow-Methods 'GET,POST,PUT,DELETE,OPTIONS' always;
   add_header Access-Control-Allow-Headers 'Authorization,X-Requested-With,Content-Type,Origin,Accept' always;

   proxy_pass http://localhost:8081/;
   }   
  ```

- 服务自身设置header头 
  
   ```golang
   w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type,Origin,X-Token")
   ```

### OPTIONS请求
[http的options请求](https://cloud.tencent.com/developer/news/397683)

#### options出现的情况

1. 获取目的资源所支持的通信方式

    黑客有可能经常用到这个；在响应报文中包含一个Allow首部字段，该字段的值表明了服务器支持的所有HTTP方法

2. 跨域请求中`options`请求是浏览器自发起的`preflight request`(预检请求)，以检测实际请求是否可以被浏览器接受

    preflight request请求报文中有两个需要关注的首部字段：
    - Access-Control-Request-Method：告知服务器实际请求所使用的HTTP方法
    - Access-Control-Request-Headers：告知服务器实际请求所携带的自定义首部字段

    服务器也会添加`origin header`,告知服务器实际请求的客户端的地址。服务器基于从预检请求获得的信息来判断，是否接受接下来的实际请求。

    服务器所返回的`Access-Control-Allow-Methods`首部字段将所有允许的请求方法告知客户端，返回将所有`Access-Control-Request-Headers`首部字段将所有允许的自定义首部字段告知客户端。此外，服务器端可返回`Access-Control-Max-Age`首部字段，允许浏览器在指定时间内，无需再发送预检请求，直接用本次结果即可。

#### 复杂的跨域请求浏览器才发送`preflight request`

- 复杂跨域请求条件：
    - 使用方法put或者delete;
    - 发送json格式的数据（content-type: application/json）
    - 请求中带有自定义头部

    不满足条件的是简单请求

#### 为什么跨域的复杂请求需要`preflight request`？

    复杂请求可能对服务器数据产生副作用。例如delete或者put,都会对服务器数据进行修改,所以在请求之前都要先询问服务器，当前网页所在域名是否在服务器的许可名单中，服务器允许后，浏览器才会发出正式的请求，否则不发送正式请求。





