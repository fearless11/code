package main

import (
	"fmt"
	"reflect"
)

// https://godoc.org/reflect
// Package reflect implements run-time reflection,
//  allowing a program to manipulate objects with arbitrary types.

// 可获取一个未知接口的数据的类型、方法、值

func main() {
	refectTypeAndValue()
	refectMethod()
}

type fruit struct {
	Name string
}

func (f *fruit) Display() {
	fmt.Println("Display:", f.Name)
}

func refectTypeAndValue() {
	fru := &fruit{Name: "apple"}
	//interface -- reflect
	ff := reflect.ValueOf(fru)
	fmt.Printf("obj: %+v\ntype: %v\nvalue: %v\n", ff, ff.Type(), ff.Elem())

	// reflect -- interface
	fu := ff.Interface().(*fruit)
	fmt.Println("name:", fu.Name)
}

func refectMethod() {
	fru := &fruit{Name: "apple"}

	obj := reflect.ValueOf(fru)
	// 只能获取到公开的方法（首字母大写）
	fmt.Println("num:", obj.NumMethod())
	for i := 0; i < obj.NumMethod(); i++ {
		fmt.Println("type: ", obj.Type().Method(i).Type)
		fmt.Println("name: ", obj.Type().Method(i).Name)
	}

	// reflect call methods 无入参设置nil, 无出参
	obj.MethodByName("Display").Call(nil)
}
