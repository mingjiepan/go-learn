package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a uint8 = 255
	//a = 299 数值溢出
	fmt.Println(a)

	//查看变量的类型
	var b = 100
	fmt.Printf("b的类型 %T \n", b)

	//怎么查看变量占用字节大小和数据类型
	fmt.Printf("b 占用字节数 %d \n", unsafe.Sizeof(b))

	//尽量使用占用空间小的数据类型
	var age int64 = 10
	fmt.Printf("age 类型 %T, 占用大小 %d", age, unsafe.Sizeof(age))
}