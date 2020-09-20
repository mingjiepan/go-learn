package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score int
}

type Pupil struct {
	Student//嵌入了student匿名结构体
}

type Graduate struct {
	Student
}

func (stu *Student) showInfo() {
	fmt.Println("name=", stu.Name, ",age=",stu.Age,",score=",stu.Score)
}

//当我们对结构体嵌入了匿名结构体后，使用方法会发生变化
func main() {
	var pupil *Pupil = &Pupil{}
	pupil.Student.Score = 20
	pupil.Student.Age = 10
	pupil.Student.Name = "lisi"
	pupil.showInfo()


	graduate := &Graduate{}
	graduate.Age = 20
	graduate.Name = "zhangsan"
	graduate.Score = 123.2
}
