package main

import "github.com/panmingjie/day01/model"

//假如结构体本身有属性A，且匿名结构体也有属性A，那么当访问结构的属性A时，采用就近原则，会访问本身的A
func main() {

	b := model.B{}
	b.Name = "adsad"

}
