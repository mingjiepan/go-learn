package main

import (
	"fmt"
	"github.com/panmingjie/day01/model"
)
//封装、继承、多态
func main() {
	var person = model.NewPerson("zhangsan")
	person.SetAge(10)
	fmt.Println(person.GetAge())
}