package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	//服务器开始监听
	listen, err := net.Listen("tcp", "localhost:8899")
	if err != nil {
		return
	}
	defer listen.Close()
	for {
		//等待客户端的连接
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("连接出错")
			break
		}
		go handleConn(accept)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr()
	fmt.Println(remoteAddr)

	for {
		fmt.Println("服务器等待客户端的输入")
		s1 := make([]byte, 10)
		read, err := conn.Read(s1)//阻塞
		if err == io.EOF {
			fmt.Println("远程客户端已退出", err)
			break
		}
		fmt.Print(string(s1[:read]))
	}
}


