package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle: 8,//最大空闲数
		MaxActive: 0,//表示和数据库的最大连接数,0表示不限制
		IdleTimeout: 100,//最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("set", "age", 10)
	if err != nil {
		fmt.Println("write data error = ", err)
		return
	}
	i, err := redis.Int(conn.Do("Get", "age"))
	if err != nil {
	}
	fmt.Println("age = ", i)
}