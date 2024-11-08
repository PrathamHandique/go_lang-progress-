package main

import (
	"fmt"
	"time"
)

func sayHello(){
	fmt.Println("Hello")
	time.Sleep(1000*time.Millisecond)
	fmt.Println("world")
}
func sayHi(){
	fmt.Println("Hi Pratham")
	time.Sleep(1000*time.Millisecond)
	//fmt.Println("pratham")
}
func main(){
	fmt.Println("main started")
	go sayHello()
	go sayHi()
	time.Sleep(2000*time.Millisecond)
}