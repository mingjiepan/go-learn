package main

import (
	"errors"
	"fmt"
)

//函数去读取

func readConfig(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("文件读取错误")
	}
}

func test2() {
	err := readConfig("hello world")
	if err != nil {
		//输出错误，并终止程序
		panic(err)
	}
	fmt.Println("test2代码继续执行")
}

func main() {
	test2()
}
