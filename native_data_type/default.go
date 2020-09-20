package main

import "fmt"

//基本数据类型的默认值
// 整型 0 浮点型 0 字符串 "" 布尔 false
//go所有的数据类型都有一个默认值

func main() {
	var a int
	var b float64
	var c string
	//%v 输出变量的默认值
	fmt.Printf("a=%v,b=%v,c=%v \n",a,b,c)

	var e = 10
	fmt.Printf("e = %v", e)
}