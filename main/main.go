package main

import (
	"fmt"
	"github.com/panmingjie/day01/utils"
	"strconv"
)

func DecToBin(n int) string {
	result := ""

	if n == 0 {
		return "0"
	}

	for ;n > 0;n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func main() {
	result := utils.Cal(1, 2)
	fmt.Println(result)

	str := "中国你好"
	bytes := []byte(str)
	fmt.Println(bytes)
	fmt.Println(len(str))
	fmt.Println(string(bytes))
	fmt.Println(DecToBin(228))
	fmt.Println(DecToBin(184))
	fmt.Println(DecToBin(173))
	fmt.Println(DecToBin(229))
}
