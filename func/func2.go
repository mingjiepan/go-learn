package main

import "fmt"

//fun1全局匿名函数
var fun1 = func(n1, n2 int) int {
	return n1 * n2
}

//匿名函数3种调用方式
func main() {

	//直接调用
	result := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(result)

	//赋值给变量，通过变量来调用
	a := func(str string) {
		fmt.Println("hello ", str)
	}
	a("zhangsan")

	result1 := fun1(2, 4)
	fmt.Println(result1)
}
