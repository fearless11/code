
## 功能

    从不同的数据源拉取数据，将数据内容与一组搜索项做对比，然后将匹配的内容显示在终端窗口

## 程序架构流程图

    主goroutine
    获取数据 ——> 执行搜索 ——> 跟踪结果 ——> 显示结果

    执行搜索的goroutine
    使用接口进行匹配 ——> 发送结果 ——> 报告任务完成

    跟踪结果的goroutine
    等待所有结果 ——> 停止工作

## 目录结构

    data
      data.json      # 包含一组数据源
    matchers
      rss.go         # 搜索rss源的匹配器
    search
      default.go     # 搜索数据用的默认匹配器
      feed.go        # 用于读取json数据文件
      match.go       # 用于支持不同匹配器的接收
      search.go      # 搜索主控制逻辑
    main.go          # 程序入口
