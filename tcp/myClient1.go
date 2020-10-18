package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main () {
	conn, err := net.Dial("tcp", "localhost:8899")
	if err != nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)

	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失败")
	}
	write, err := conn.Write([]byte(readString))
	if err != nil {
	}
	fmt.Println("write length", write)
}
