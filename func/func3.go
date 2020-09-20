package main

import "fmt"

//闭包就是一个函数和与其相关的引用环境组合的一个整体（实体）
//一个函数A的返回值是一个函数类型，返回值是定义在函数A内部的一个匿名函数，匿名函数访问定义在函数A的变量
//返回的匿名函数和匿名函数用到的函数外的变量，共同构成一个闭包

//累加器
func accumulate() func(int)int {
	var n int = 0
	return func(a int) int {
		n = n + a
		return n
	}
}

func main() {
	fun1 := accumulate()
	fun1(1)
	fun1(2)
	result := fun1(3)
	fmt.Println(result)
}