package main

import (
	"fmt"
	"strconv"
)
func main(){
	fmt.Println("we will learn about data conversion")
	var num1 int = 10
	str :=strconv.Itoa(num1)
	fmt.Println("the string is",str)
	fmt.Printf("the type of str is %T\n",str)

	num_string :="12.3"
	num_float,_ :=strconv.ParseFloat(num_string,64)
	fmt.Println("the float is",num_float)
	fmt.Printf("the type of num_float is %T\n",num_float)

}