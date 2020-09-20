package main

import "fmt"

func main() {
	var arr1 [5]int
	arr1[0] = 1
	arr1[1] = 2

	//数组的地址，就是数组第一个元素的地址,，，数组的第二个元素的地址，就是第一个元素地址加上每个元素占用的字节数，对int来说，加8个字节
	//所以，通过数组，便可以拿到第n个位置的值，因为地址的内存分布是连续的，且固定大小的
	fmt.Printf("arr1 = %p, arr1[0] = %p, arr1[1] = %p", &arr1, &arr1[0], &arr1[1])

	//数组的其他声明方式
	var arr2 [5]int = [5]int {}
	fmt.Println(arr2)

	var arr3 = [5]int {1}
	fmt.Println(arr3)
	arr4 := [...]int {}
	fmt.Println(arr4)
	arr5 := [3]int {1:2}
	fmt.Println(arr5)
}
