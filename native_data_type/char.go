package main

import "fmt"

//go的字符
func main() {
	//输出byte值，是直接输出Ascii码值
	var a byte = 'a'
	fmt.Println(a)

	var b byte = '0'
	fmt.Println(b)

	//如果要输出字符，需要格式化输出
	fmt.Printf("a=%c b=%c \n", a, b)

	//unicode码值21217大于255了，所以不能用byte来存储
	var c int16 = '北'
	fmt.Println(c)
	fmt.Printf("c = %c \n", c)

	//字符可以参与运算
	d := '北'
	e := '京'
	fmt.Printf("d + e = %c\n", d + e)

	fmt.Println('a' + 10)
}