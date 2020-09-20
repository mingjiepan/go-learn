package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 100
	var b float64 = 12.12
	var c bool = true
	var d byte = 'N'

	var str string

	str = strconv.FormatInt(int64(a), 10)
	fmt.Printf("str type %T str=%v \n", str, str)

	//f
	str = strconv.FormatFloat(b,'f', 10, 64)
	fmt.Printf("str type %T str=%v \n", str, str)

	str = fmt.Sprintf("%t", c)
	fmt.Printf("str type %T str=%v \n", str, str)

	str = fmt.Sprintf("%c", d)
	fmt.Printf("str type %T str=%v \n", str, str)
}
