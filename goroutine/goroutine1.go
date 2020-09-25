package main

import (
	"fmt"
	"strconv"
	"time"
)

/*协程是轻量级线程（编译器优化）
有独立的栈空间
共享程序堆空间
调度由用户控制*/

func helloworld() {
	for i:=0; i< 10; i++ {
		fmt.Println("hello world", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go helloworld()//开启一个协程

	for i:=0; i< 10; i++ {
		fmt.Println("hello golang", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}