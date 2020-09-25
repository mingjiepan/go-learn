package main

import (
	"fmt"
	"sync"
	"time"
)

//go build -race test3.go
//编写一个函数，计算1-200各个数的阶乘，并放入map中
//我们启动的协程多个，并放入map中

//channel管道的需求
//全局变量加入互斥锁
var (
	myMap = make(map[int]int, 10)
	//声明一个全局互斥锁
	lock sync.Mutex
)

//计算某个数的阶乘，然后将这个结果放入到map中
func test(n int) {
	var res int = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//出现并发写map问题
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 10; i++ {
		go test(i)
	}

	//目前通过睡眠来处理
	time.Sleep(time.Second * 3)

	//这边也需要进行上锁，虽然程序在sleep之后，就不存在写的竞争了，但是底层编译器并不知道程序在上面睡眠了，
	lock.Lock()
	for key,value := range myMap {
		fmt.Printf("map[%d]=%d\n", key, value)
	}
	lock.Unlock()
}