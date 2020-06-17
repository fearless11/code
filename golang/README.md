



#### Tips

##### Markdown目录生成方式
```bash
# 全局安装doctoc插件 i: install g:global
npm i doctoc -g   

# 文档生成目录
doctoc README.md  
```



#### make & new

- make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。
- new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。
- make(T, args)初始化了内部的数据结构，填充适当的值，并且返回一个有初始值(非零)的T类型。

  ```bash
  	m := new(map[string]string)
	m1 := make(map[string]string)
	fmt.Println(*m == nil, m1 == nil)
    #output: true false
  ```