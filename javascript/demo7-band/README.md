
## 创建一个网站

### 设计原型效果

### 页面结构（结构层）

- 头部区域。包含站点的品牌Logo信息，用<header>元素
- 导航区域。放一组链接指向各个页面，用<nav>元素
- 内容区域。每一页包含实质性内容，用<article>元素

### CSS（样式层）

- 颜色
- 布局： 一栏或多栏 
- 版式： 字体和大小

### JavaScript（行为层）

- 高亮: 结合Class属性+CSS实现
- 内部导航隐藏: 结合style对象实现
    ```HTML
    <!-- style是一个对象 -->
    <p id="example" style="color: grey; font-family: 'Arial', sans-serif;">an example of a paragraph</p>
    ```

    ```JavaScript
    <script>
    var para = document.getElementById("example");
    alert(typeof para.style);
    alert("this color is" + para.style.color);
    alert("this font family is" + para.style.fontFamily);
    // DOM根据style设置样式
    para.style.backgroundColor="#ffc"
    </script>
    ```
  
- 异步请求: Ajax结合提交事件实现
