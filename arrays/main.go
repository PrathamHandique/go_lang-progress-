package main

import "fmt"

func main() {
	fmt.Println("learning arrays")
	var name[5] string
	name[0]="Prince"
	name[1]="Kumar"
	fmt.Println("the names of array are ",name)
	var numbers = [5] int {1,2,3,4,5} //another way to initialize array
	fmt.Println("the numbers of array are ",numbers)
	fmt.Println("the length of array is ",len(numbers))
	fmt.Println("the first element of array is ",numbers[0])
	fmt.Printf("Price array is %q\n",name) //%q is used to print the array , here q is for quoted string

}