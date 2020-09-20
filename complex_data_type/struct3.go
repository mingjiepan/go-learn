package main

import "fmt"

type struct1 struct {
	name string
}
type struct2 struct {
	name string
}

type ss struct1

//结构体之间的类型转换，需要两种类型的结构体字段完全一致，才能转换
//结构体进行type重新定义（相当于取别名），go认为是新的数据类型，但是相互之间可以进行转换
func main() {
	var s1 struct1 = struct1{"hello"}
	var s2 struct2 = struct2(s1)
	var s3 ss = ss(s2)
	fmt.Println(s2)
	fmt.Println(s3)
}