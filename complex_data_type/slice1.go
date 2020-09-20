package main

import "fmt"

//切片是数组的引用，长度可以变化，是一个动态数组，声明一个切片时，和声明数组一样，除了不需要写数组长度
//var s1 []int
func main() {
	var arr1 [5]int = [5]int {11, 22, 33}
	slice1 := arr1[1:3]
	slice1[0] = 111
	fmt.Println(arr1)
}