package main

import "fmt"

func main() {
	var a int = 100
	var b float64 = 12.12
	var c bool = true
	var d byte = 'N'

	var str string

	str = fmt.Sprintf("%d", a)
	fmt.Printf("str type %T str=%v \n", str, str)

	str = fmt.Sprintf("%f", b)
	fmt.Printf("str type %T str=%v \n", str, str)

	str = fmt.Sprintf("%t", c)
	fmt.Printf("str type %T str=%v \n", str, str)

	str = fmt.Sprintf("%c", d)
	fmt.Printf("str type %T str=%v \n", str, str)
}