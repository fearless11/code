package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const dbtmpl = `
<h1>grocer shopping</h1>
<table>
<tr style='text-align: left'>
  <th>name</th>
  <th>price</th>
</tr>
{{ range $key,$value := .}}
<tr>
  <td>{{$key}}</a></td>
  <td>{{$value}}</a></td>
</tr>
{{end}}
</table>
`

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	// http.HandlerFunc(db.list)
	// 是类型转换，不是函数调用
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	grocer := template.Must(template.New("grocer").Parse(dbtmpl))
	if err := grocer.Execute(w, db); err != nil {
		fmt.Fprintf(w, err.Error())
	}
	// for item, price := range db {
	// 	fmt.Fprintf(w, "%s: %s\n", item, price)
	// }
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item:%q \n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
