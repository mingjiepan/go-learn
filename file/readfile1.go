package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("F:\\usr\\tools\\go\\src\\github.com\\mingjiepan\\go-learn\\complex_data_type\\struct1.go")
	if err != nil {
		fmt.Println("读取文件失败")
	}
	newReader := bufio.NewReader(file)
	for {
		readLine, _, err := newReader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(readLine))
	}
}
