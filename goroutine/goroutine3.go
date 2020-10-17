package main

import "fmt"

func writeData(intChan chan int) {
	for i := 0; i < 5000; i++ {
		intChan <- i
	}
	close(intChan)
}

func readData(intChan chan int, booChan chan bool) {
	for {
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Println(v)
	}
	booChan <- true
	close(booChan)
}

func main() {
	intChan := make(chan int, 5000)
	boolChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, boolChan)
	<- boolChan
}
