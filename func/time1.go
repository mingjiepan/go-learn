package main

import (
	"fmt"
	"time"
)

//时间和日期相关函数
func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now)

	fmt.Printf("now的值%v \n now的类型%T \n", now, now)
	fmt.Printf("year = %v, month = %v day = %v", now.Year(), int(now.Month()), now.Day())
}
