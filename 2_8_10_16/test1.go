package main

import "fmt"


//go中，没法直接用二进制
func main() {

	var a int = 10
	fmt.Printf("a 的二进制 %b", a)


	var b int = 011//这里是八进制
	fmt.Println("b的十进制", b)

	var c int = 0x11
	fmt.Println("c的十进制", c)

}
