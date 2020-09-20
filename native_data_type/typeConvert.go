package main

import "fmt"

//go不能类型的变量之间赋值时需要显示转换，不能自动转换
//表达式 T(v)将值转换为类型 T
func main()  {
	var a int = 10
	var b float64 = float64(a)
	fmt.Println(b)

	//将范围大的值转换成范围小的，不会报错，但是结果不对
	fmt.Println("------")
	var c int64 = 10000
	fmt.Println(int8(c))
}