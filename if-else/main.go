package main

import "fmt"
func main(){
	x:=4
	y:=10
	if x <5 && (y>5 || y<43){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
}