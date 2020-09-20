package model


type Student struct {
	Name string
	Age int
}

type Point struct {
	X int
	Y int
}

type Rect struct {
	Left, Right Point
	Up *Point
}

type A struct {
	Name string
	age int//注意，这边的age是小写
	Address string
}

type B struct {
	A
	Name string
}