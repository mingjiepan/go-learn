package main

import "fmt"

type BaseOpus struct {
	opusName string
	opusId int
}

type Comic struct {
	BaseOpus
	imgUrl string
}

type FunShoot struct {
	BaseOpus
	videoUrl string
	status int
}

//假如结构体本身有属性A，且匿名结构体也有属性A，那么当访问结构的属性A时，采用就近原则，会访问本身的A
func main() {
	comic := Comic{BaseOpus:BaseOpus{opusName: "葫芦娃", opusId: 123}, imgUrl: "http://cn.migudm//123.png"}
	fmt.Println("opusName=", comic.opusName, ", opusId=",comic.opusId)

	var funShoot FunShoot
	funShoot.BaseOpus.opusName = "七剑下天山"
	funShoot.opusId = 234
	funShoot.status = 1
	fmt.Println("funShoot.opusName", funShoot.opusName, ", funShoot.opusId=", funShoot.opusId)
}
