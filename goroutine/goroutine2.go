package main

import (
	"fmt"
	"runtime"
)

/*调度模型 MPG
M：操作系统的主线程（物理线程）
P：协程执行需要的上下文
G：协程*/
func main() {
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
	gomaxprocs := runtime.GOMAXPROCS(numCPU - 1)
	fmt.Println(gomaxprocs)
}