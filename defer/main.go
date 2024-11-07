package main

import "fmt"
func add(a,b int)int{
	return a+b
}
func main(){
	fmt.Println("start of the program")
	data:=add(10,20)
	defer fmt.Println("the sum is",data)
	defer fmt.Println("middle of the program")
	fmt.Println("end of the program")
}