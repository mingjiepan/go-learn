package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

//使用tag标签，用于在序列化成json时使用，不能直接将结构体的字段名称改为小写，因为这样在json包就访问不了结构体的私有属性了
func main() {
	var s1 Teacher = Teacher{Name: "张三", Age: 10}
	jsonStr,err := json.Marshal(s1)
	if err != nil {
		fmt.Println("发生错误")
	}
	fmt.Println(string(jsonStr))
}
