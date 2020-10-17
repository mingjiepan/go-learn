package main

import "fmt"

//使用close来关闭管道，当channel关闭后，就不能再向channel写入数据了，但是，可以读取数据
//在遍历管道时，如果channel没有关闭，那么遍历完元素后会产生deadlock
//如果channel关闭了，那么可以正常遍历完管道
func main() {

	intChan := make(chan int, 5)
	intChan <- 10
	intChan <- 20
	intChan <- 30
	intChan <- 40
	intChan <- 50

	close(intChan)

	for value := range intChan {
		fmt.Println(value)
	}
}
