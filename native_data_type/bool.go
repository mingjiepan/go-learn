package main

import (
	"fmt"
	"unsafe"
)

//布尔类型占用1个字节
func main() {
	var a bool = true
	var b bool = false

	fmt.Println("a 的存储空间", unsafe.Sizeof(a))
	fmt.Println("b 的存储空间", unsafe.Sizeof(b))
}
