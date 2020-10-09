package main

import "fmt"

//接口包含一组方法，但是不能包含任何变量
//实现了某个接口，需要实现所有的方法，golang中不需要使用implements关键字，，也因为这个特性，所以go中耦合度很低

//接口本身不能创建实例
//只要是自定义类型，就都可以实现接口，不仅仅是针对结构体
//自定义类型可以实现多个接口

//接口数组中，可以存放实现了该接口的任何类型实例

//Integer类型也实现了StudentService
type Integer int
func (i Integer) addOne() int {
	return int(i + 1)
}

func main() {
	var i Integer = 2
	result := i.addOne()
	fmt.Println(result)
}