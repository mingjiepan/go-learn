package main

import "fmt"

func main() {
	//基本数据类型在内存的布局
	var i int = 10
	fmt.Println("i 的地址", &i)

	//指针变量：存储地址的变量
	var p1 *int = &i
	fmt.Println("p1 本身的地址", &p1)
	fmt.Println("p1 指向的地址", p1)
	fmt.Println("p1 指向地址所存储的值", *p1)
}