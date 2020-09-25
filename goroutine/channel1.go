package main

import "fmt"

//channel本身是一个队列（先进先出），线程安全的，多gotoutine访问时，无需加锁
//channel是引用类型的，必须初始化后才能使用，即make之后
//channel是有类型的，，声明方式  var 变量名 chan 数据类型

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan 的值=%v\n", intChan)
}