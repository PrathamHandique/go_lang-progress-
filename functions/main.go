package main

import (
	"fmt"
)

func simpleFunction(){
	fmt.Println("This is a simple function")
}
func add(a, b int) int {  // also fine if we  write add(a int, b int) int
	return a+b
} 

func sub(a, b int) (result int) { // also fine if we write sub(a int, b int) (result int)
	result = a-b
	return
}
func main(){
	fmt.Println("learning funcitons")
	simpleFunction()
	ans := add(2,3)
	fmt.Println("The sum is ", ans)
	and :=sub(5,3)

	fmt.Println("The difference is ", and)
}