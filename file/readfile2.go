package main

import (
	"fmt"
	"io/ioutil"
)

//适合读取文件不大的情况
func main() {
	file, err := ioutil.ReadFile("/Users/panmingjie/Develop/go/src/github.com/panmingjie/day01/object/test3.go")
	if err != nil {
		fmt.Println("读取文件错误")
	}
	fmt.Println(string(file))
}
