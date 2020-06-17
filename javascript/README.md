
<!-- TOC -->

- [JavaScript DOM 编程艺术 (第二版) 学习笔记](#javascript-dom-编程艺术-第二版-学习笔记)
    - [基本术语](#基本术语)
    - [简介语法](#简介语法)
        - [HTML](#html)
        - [CSS](#css)
        - [JavaScript](#javascript)
    - [最佳实践](#最佳实践)
    - [实用网站](#实用网站)

<!-- /TOC -->

## JavaScript DOM 编程艺术 (第二版) 学习笔记

### 基本术语

- W3C：万维网联盟组织。

- ECMAScript：Netspace和Sum公司联合ECMA（欧洲计算机制造商协会）对JavaScript语言进行了标准化后出现的名字。

- JavaScript：由JavaScript与Sun公司开发的Java程序语言没有任何联系，命名大概是想听起来系出名门的感觉。脚本语言，由Web浏览器进行解释和执行。

- DOM：文档对象模型。一套对文档的内容进行抽象和概念化的方法。一种适用于多种环境和多种程序设计语言的通用型API。

- BOM：浏览器对象模型。一套对浏览器的操作API。

- HTML5： 一种HTML语言当前及未来的新标准。

- Ajax： 2005年Adaptive Path公司的Jesse James Garret发明。核心是XMLHttpRequest，依赖JavaScript。用于异步加载页面内容的一种技术。（具体过程—浏览器发出请求后触发事件消息被Ajax拦截，代理浏览器与服务端进行操作。优点—减少重复加载整个页面的次数，实现局部更新。）
- 解释型&编译型

  - 解释型：不需要编译器；代码错误执行时发现；速度慢可移植性差；学习成本低  

  - 编译型：需要编译器编译为执行文件；代码错误编译阶段发现；速度快移植性高；学习成本高

- 强语言&弱语言

  - 强语言： 必须明确类型声明，之后变量只能存放声明类型 `int型变量只能存整数`

  - 弱语言：可以在任何阶段改变变量存放的数据类型 `int型变量可以存放字符串`

- Web设计的网页分为三层

  - 结构层（HTML、XHTML、HTML5）
  
  - 样式层（CSS）
  
  - 行为层（事件响应DOM脚本+JavaScript）

### 简介语法

#### HTML

```HTML
<!-- 声明为 HTML5 文档 , 标签有书写规则要求 -->
<!DOCTYPE html> 

<!-- 标签 等价于 元素 -->
<b>成对出现</b>

<!-- 元素带属性 -->
<form action="demo-form.php" autocomplete="on"></form>
```

#### CSS

- HTML引入CSS

```HTML
<!-- 引入CSS -->
<head>
    <link rel="stylesheet" media="screen" href="styles/basic.css"/>
</head>
```
- basic.css

```css
/* 分层维护样式，导入到基本样式 */
@import url(layout.css);
@import url(color.css);
@import url(typography.css);
```

- layout.css

```css
selector {
    property: value;
}

/* 通配元素 */
* {
    padding: 0;
    margin: 0;
}

/* 单个元素 */ 
body {
    color: white;
    background-color: block;
}
header nav li a {
    display: block;
    float: left;
    padding: .5em 2em;
    border-right: .1em solid;
}

/*
 class属性 区分一个元素与其他元素
 <p class="special">This is special class</p>
 <h2 class="special">h2 special class</h2> 
*/
.special {
    font-style: italic;
}
h2.special {
    text-transform: uppercase;
}

/*
 id属性 区分单个元素
 <ul id="purchases"><li></li></ul>
*/
#purchases{
    border: 1px solid white;
}
#purchases li {
    font-weight: bold;
}
```

#### JavaScript

- HTML导入JavaScript

    ```HTML
    <!-- 第一种： 放在<head>标签中 -->
    <head> 
      <script> xxx </script>
    </head>

    <!-- 第二种：存为.js文件，在<script>标签中加src属性 -->
    <head> 
       <script src="file.js"></script>
    </head>

    <!--  第三种：把<script>标签放在HTML文档最后</body>标签之前 -->
    <body>
       <script src="file.js"></script>
    </body>
    ```

- 工具

  浏览器F12进入到Consle,就可以练习


- 注释

  ```JavaScript
  // 单行 
  /* 多行 */
  ```

- 变量

  ```JavaScript
  // 弱数据类型
  my_mood="sad"
  var mood="happy", age=33;  
  ```

- 数据类型

  字符串

  ```javascript
  // 单双引号 
  mood="happy" ;
  mood='happy' ;
  ```

  数值

  ```javascript
  // 浮点数负数 
  age=20.1;
  age=-20;
  ```

  布尔值

  ```javascript
  // true/false 
  sleeping=true
  ```

  数组

  ```javascript
  var beatles=Array(4);
  var beatles=Array();
  beatles[index]=value;
  
  var beatles=Array("Mary","John","Paul");
  var beatles=["Mary","John","Paul"];
  ```

  关联数组

  ```javascript
  var beatles=Array();
  beatles['name']='John';
  beatles['year']=1940;
  ```

  对象

  ```javascript
  // 自定义对象
  var lennon=Object();
  lennon.name="John";
  lennon.year=1940;
  
  var lennon={};
  var lennon={name:"John", year:1940};
  
  // 内建对象
  var beatles = new Array();
  beatles.length 
  var date = new Date();
  var today = date.getDay();
  
  // 宿主对象
  // 由浏览器提供的对象
  ```

- 操作

  拼接

  ```javascript
  // 字符串 + 数值 = 字符串 （弱类型）
  // 数值 + 数值 = 数值 
  var mood = "happy";
  var message = "I am feeling " + mood;
  
  var year = 2005;
  var message = "This year is " + year;
  
  alert("10"+20)
  alert(10+20)
  ```

  判等

  ```javascript
  // == 非严格
  // === 严格
  
  if (a === b) {
      alert("a qeuals b");
  }
  ```

  循环

  ```javascript
  while (condition) {
    statement;
  }
  
  for (var count=1; count < 1; count++ ) {
      statement;
  }
  ```

  DOM
  ```javascript
  // 获取body元素个数
  alert(document.getElementsByTagName("body").length)
  ```

### 最佳实践 

- 平稳退化

  浏览器不支持JavaScript仍然能访问网站。

- 渐进增强

  能够用Ajax实现的应用一定也可以通过非Ajax技术实现。考虑网站的可用性和可访问性。一开始设计时是常规网站，然后用Ajax增强功能。而不是一开始设计就以Ajax为起点。

- 分离JavaScript与CSS

  把网页的结构与JavaScript脚本动作分开。

- 向后兼容

  确保老版本浏览器不会因JavaScript而死掉。

- 性能考虑

  - 尽可能少访问DOM
  - 合并脚本和存访位置
  - 压缩脚本：把脚本中不必要的字节，如空格和注释统统删除；短变量替换

### 实用网站

- [动画加载GIF-ajaxload.info](http://ajaxload.info/)
- [在线压缩JS/CSS代码](https://tool.oschina.net/jscompress/)
- [菜鸟教程-runoob.com](https://www.runoob.com/html/html-tutorial.html)
- [HTML5特征-modernizr.com](https://modernizr.com/download?setclasses)