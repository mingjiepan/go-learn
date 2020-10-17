package main

import "fmt"

type People struct {
	name string
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 5)

	people := People{name: "hello"}

	allChan <- 10
	allChan <- "hello world"
	allChan <- people
	allChan <- true

	myMap := make(map[int]string, 10)
	myMap[0] = "zhangsan"
	allChan <- myMap

	//要获取到people对象，需要将前两个元素推出
	<-allChan
	<-allChan

	mypeople := <- allChan
	fmt.Printf("mypeople的类型=%T", mypeople)

	//取出的类型是interface，需要转换为实际的类型
	p := mypeople.(People)
	fmt.Println(p.name)
}