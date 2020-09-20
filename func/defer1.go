package main

import "fmt"

//当执行到defer时，暂不执行，会将defer后面的语句压入到独立的栈（defer栈）
//当函数执行完后，再从defer栈，按照后入先出的方式出栈，执行

//defer的用法，通常是在创建好资源后，马山调用defer 资源关闭，防止后面忘记关闭相应的资源了
func sum(n1, n2 int) int {
	defer fmt.Println("n1 = ", n1)
	defer fmt.Println("n2 = ", n2)
	res := n1 + n2
	fmt.Println("res = ", res)
	return res
}

func main() {
	res := sum(1, 2)
	fmt.Println(res)
}