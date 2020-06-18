package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// 序列化、反序列化、格式化输出

// https://godoc.org/encoding/json
// JSON: 轻量级的文本交换格式. 类似 XML、ini、yaml、toml
// 规范: key/value
//  key: string
//  value: 数字、字符串、数组 [...]、对象 {...}、null
//  tag: `json:"name"`

type student struct {
	Name   string `json:"name"`           // 用tag标记key为name
	City   string `json:"city,omitempty"` // key为city，字段为零值将不显示字段
	Age    int    `json:",omitempty"`     // key为Age，字段为零值将不显示字段
	School string `json:"-"`              // 忽略该字段
	Hobby  string `json:"-,"`             // key为-
}

// marshal 序列化
// object --> byte (内存 ——> 磁盘/网络) 把对象转换成字节序列的过程
func marshal() {
	stu := student{
		Name:   "Tom",
		City:   "ShenZhen",
		Age:    20,
		School: "Yale",
		Hobby:  "Music",
	}
	b, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	//  格式化
	s, err := json.MarshalIndent(stu, " ", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(s))
}

// umarshal 反序列化
// byte --> object (磁盘/网络 --> 内存) 把字节对象恢复会对象的过程
/*
 bool    --- JSON booleans
 float64 --- JSON numbers
 string  --- JSON strings
 []interface{} --- JSON arrays
 map[string]interface{}, --- JSON objects
 nil --- null
*/
func umarshal() {
	byt := []byte(`{
		"Name":"Tom",
		"City":"ShenZhen"
		}`)

	// map[string]interface{} for JSON objects
	var stu map[string]interface{}
	err := json.Unmarshal(byt, &stu)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(stu)

	// []interface{} for JSON arrays
	stb := []byte(`["aa","bb","cc"]`)
	var arr []interface{}
	err = json.Unmarshal(stb, &arr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(arr)
}

// decode解码 stream --> json
func decode() {
	const streamJSON = `[
		{"Name":"Tom"},
		{"Name":"Marry"}
	]`

	dec := json.NewDecoder(strings.NewReader(streamJSON))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		fmt.Println(t, err)
		return
	}

	// while the array contains values
	for dec.More() {
		pp := struct{ Name string }{}
		err := dec.Decode(&pp)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("name:", pp.Name)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		fmt.Println(t, err)
		return
	}
}

func main() {
	marshal()
	umarshal()
	decode()
}
