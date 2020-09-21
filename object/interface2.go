package main

import "fmt"

//interface类型默认是一个指针（引用类型），如果没有对interface初始化就使用，那么会输出nil
//空接口interface{}没有任何方法，所以所有类型都实现了空接口
func method1(i interface{}) {
	fmt.Println(i)
}

func main() {
	method1(1)
	method1("hello world")
	method1(true)
}