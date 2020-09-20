package main

import "fmt"
import "github.com/panmingjie/day01/model"

//声明结构体的4种方式,,2种是返回结构体变量，2种是返回结构体指针
func main() {
	var s1 model.Student
	fmt.Println(s1)

	var s2 model.Student = model.Student{}
	fmt.Println(s2)

	var s3 *model.Student = new(model.Student)
	(*s3).Age = 10
	s3.Name = "zhangsan"
	fmt.Println(s3)

	var s4 *model.Student = &model.Student{}
	(*s4).Age = 20//标准访问方式
	s4.Name = "lisi"//简化方式,go的底层会转换为(*s4)
	fmt.Println(s4)
}

