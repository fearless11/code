

#### LSP里氏转换
  可替换型（LSP里氏替换）：一个类型可以自由地被另一个满足相同接口的类型替换。
  满足同一接口的不同类型是可替换的。

- HandlerFunc实现了Handler的接口
```
package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)    
}
```
- 转换
  http.HandlerFunc(db.list) 是显示转换类型，不是函数调用
```
func main() {
    db := database{"shoes": 50, "socks": 5}
    mux := http.NewServeMux()
    // 里氏转换
    mux.Handle("/list", http.HandlerFunc(db.list))
   log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
    for item, price := range db {
        fmt.Fprintf(w, "%s: %s\n", item, price)
    }
}
```

