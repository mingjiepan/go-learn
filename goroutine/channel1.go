package main

import "fmt"

//channel本身是一个队列（先进先出），线程安全的，多gotoutine访问时，无需加锁
//channel是引用类型的，必须初始化后才能使用，即make之后
//channel是有类型的，，声明方式  var 变量名 chan 数据类型

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan 的值=%v\n", intChan)
	fmt.Println("channel 本身的地址=", &intChan)

	//向管道写入数时，不能超过容量
	intChan <- 10
	num := 988
	intChan <- num
	//intChan <- 211

	fmt.Println("管道的长度=",len(intChan),",容量=", cap(intChan))

	//从管道读取数据
	result := <- intChan
	fmt.Println(result)
	result = <- intChan
	fmt.Println(result)

	//在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报死锁
}