
<!-- TOC -->

- [进阶](#进阶)
    - [术语](#术语)
    - [核心插件](#核心插件)
        - [Vue Router 管理组件与路由的映射关系](#vue-router-管理组件与路由的映射关系)
        - [vuex 存储组件的共享状态](#vuex-存储组件的共享状态)
- [框架](#框架)
- [文档](#文档)

<!-- /TOC -->

## 进阶

### 术语
- 实例、组件、属性、事件、生命周期、钩子函数
-  vue-cli：（命令）创建项目的脚手架
- vue-devtools：（工具）浏览器调试
-  vue router：（插件）路由与组件的映射
-  vuex： （插件）集中存储多个组件中的共享状态

### 核心插件

#### Vue Router 管理组件与路由的映射关系
<!-- 理解：根据url不同加载vue的不同组件，把nginx的路由转到内部 -->

安装
```bash
# 装插件
npm install vue-router

# 代码里装
import Vue from 'vue'
import VueRouter from 'vue-router'
# 安装插件
Vue.use(VueRouter)
```

基础

- 静态路由：路由与组件对应
- 动态路由：多个路由对应一个组件
- 嵌套路由：对路由url提供分层
- 路由重定向、命名、别名： redirect、name、alias
- 编程式的导航： 与`<router-link>`不同，通过事件或钩子触发路由跳转
- 命名视图： 一个组件对应一个视图，而不是多个组件嵌套到一个视图
- 路由组件传参：props

进阶
- 路由懒加载：不同路由对应的组件分割成不同的代码块，然后当路由被访问的时候才加载对应组件。


#### vuex 存储组件的共享状态
<!-- 类似进程间通信，把不同组件中的共享数据进行管理。可以单独定义一个store来实现 -->

安装

```bash
# vue-cli创建项目
npm install -g @vue/cli
vue create vue-vuex
cd vue-vuex
# 装插件，对在store/index.js操作
vue add vuex
```

why

- 状态管理： state（要共享的数据源）、 view（将state数据映射到视图）、actions（共享数据操作后view的变化）
- 解决问题：多个视图依赖同一状态；不同视图行为变更同一状态

核心

- state： 共享数据
- getter：需要多个地方都用的计算属性可以合并过来`this.$store.getters.result`
- mutation： 由事件类型和回调函数组成，同步操作。更新对象`Vue.set(obj, 'newProp', 123)` 或者 `state.obj = { ...state.obj, newProp: 123 }`
- action：提交mutation包含异步操作
- module：分层操作更加清晰

项目结构

1. 应用层级的状态应该集中到单个 store 对象中。
2. 提交 mutation 是更改状态的唯一方法，并且这个过程是同步的。
3. 异步逻辑都应该封装到 action 里面。
    ```shell
    ├── index.html
    ├── main.js
    ├── api
    │   └── ... # 抽取出API请求
    ├── components
    │   ├── App.vue
    │   └── ...
    └── store
        ├── index.js          # 我们组装模块并导出 store 的地方
        ├── actions.js        # 根级别的 action
        ├── mutations.js      # 根级别的 mutation
        └── modules
            ├── cart.js       # 购物车模块
            └── products.js   # 产品模块
    ```


## 框架

- [在线体验](https://panjiachen.github.io/vue-element-admin/#/login?redirect=%2Fdashboard)

- 基础模板 
  ```bash
  git clone https://github.com/PanJiaChen/vue-admin-template.git myapp
  cd myapp
  # 安装依赖
  npm install --verbose --registry=https://registry.npm.taobao.org
  # 开发环境运行
  npm run dev
  
  # 调优
  npm run preview
  # 调优查看
  npm run preview -- --report
  # 代码格式
  npm run lint
  # 检查代码格式
  npm run lint -- --fix
  # 生产dist文件
  npm run build:prod 
  # 本地预览
  npm install -g serve --verbose
  serve -s dist
  ```

- 理解

  vue核心、vue-router管理路由、vuex管理状态、Element-UI设置模板组合而成。

  ```bash
  main.js -> App.vue -> layout/index.vue (由router决定，读取store里的信息) -> layout/components/sidebar/index.vue 、layout/components/Navbar.vue 、layout/components/AppMain.vue 
  ```

- 开发设计功能点

  - 搜索
  - 排序
  - 加载
  - 分页——插件实现
  - 表单 
  - 侧边栏Logo


## 文档

[vuejs.org-guide](https://cn.vuejs.org/v2/guide/installation.html)

[vue-cli](https://cli.vuejs.org/zh/)

[vue-router](https://router.vuejs.org/zh/guide/)

[vuex-状态管理](https://vuex.vuejs.org/zh/installation.html)

[vetur-vscode插件](https://vuejs.github.io/vetur/setup.html#extensions)

[vue-elemnt-admin-框架](https://panjiachen.github.io/vue-element-admin-site/zh/guide/)

[Element-UI](https://element.eleme.cn/#/zh-CN/component/installation)
