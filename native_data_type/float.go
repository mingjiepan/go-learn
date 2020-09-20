package main

import "fmt"

func main() {
	//float32精度损失，golang默认的浮点类型为float64
	var a float32 = -123.0000091
	var b float64 = -123.0000091
	fmt.Println("a=",a, ",b=", b)

	//省略小数点前面的0
	n1 := .123
	fmt.Println(n1)

	n2 := 5.123e2
	fmt.Println(n2)

	// 10000 / (10*10)
	n3 := 10000e-2
	fmt.Println(n3)
}