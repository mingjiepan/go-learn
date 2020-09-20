package main

import (
	"fmt"
	"github.com/panmingjie/day01/model"
)


func main() {
	var s = &model.Student{}
	fmt.Printf("s指向的地址 %p \n", s)
	fmt.Printf("s本身的地址 %p \n", &s)

	//结构体字段在内存的分布是连续的，因此在取字段时，速度比较快，是基于地址+-字节数在取值的
	var react model.Rect = model.Rect{model.Point{1, 2}, model.Point{3, 4}, &model.Point{10, 20}}
	fmt.Printf("left.X=%p, left.Y=%p, right.X=%p, right.Y=%p\n", &react.Left.X, &react.Left.Y,
		&react.Right.X, &react.Right.Y)

	//up本身的地址在结构体内也是连续的，但是指向的地址不一定是连续的，通常情况下是不连续的
	fmt.Printf("up本身的地址=%p, up指向的地址=%p\n", &react.Up, react.Up)
	//left.X=0xc0000aa020, left.Y=0xc0000aa028, right.X=0xc0000aa030, right.Y=0xc0000aa038
}