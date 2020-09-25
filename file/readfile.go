package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/Users/panmingjie/Develop/go/src/github.com/panmingjie/day01/object/test3.go")
	if err != nil {
		panic("打开文件错误")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		fmt.Println(readString)
	}
}
