package model

type Dog struct {
	Name string
}

type person struct {
	name string
	age int
	salary float64
}
//构造方法
func NewPerson(name string) *person {
	return &person{name: name}
}
//设置set方法
func (p *person) SetAge(age int) {
	p.age = age
}
//get方法
func (p *person)GetAge() int {
	return p.age
}
