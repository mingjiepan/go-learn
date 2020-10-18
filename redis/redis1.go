package main

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		errors.New("链接到Redis出错")
	}
	defer conn.Close()

	do, err := conn.Do("set", "name", "zhangsan")
	if err != nil {
		errors.New("set 出错")
	}
	fmt.Println(do)

	reply, err := conn.Do("get", "name")
	fmt.Println(string(reply.([]uint8)))
	fmt.Println(redis.String(reply, err))
}