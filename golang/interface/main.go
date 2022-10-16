package main

import "fmt"

type MyError struct {
}

func (m MyError) Error() string {
	return "my error"
}

func GetMyError() error {
	var p *MyError = nil
	return p
}

func main() {
	err := GetMyError()
	fmt.Println(err == nil)
}
