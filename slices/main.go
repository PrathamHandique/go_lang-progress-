package main

import "fmt"
func main(){
	fmt.Println("learning slices")
	//slices are like dynamic arrays with no fixed size , size can be increased or decreased
	//slices are reference types , which means they are passed by reference
	//slices are like views on arrays
	// numbers := []int{1,2,3,4,5}
	// numbers=append(numbers,6,7,8,9,10)
	// fmt.Println("the numbers of slice are ",numbers)
	// fmt.Printf("the datatype of slice is %T\n",numbers)
	// name :=[] string{}
	// fmt.Println("the names of slice are ",name)	
	numbers := make ([]int ,3,5) //here 3 is length and 5 is capacity
	numbers = append(numbers,4)
	numbers = append(numbers,5)
	numbers = append(numbers,6)
	fmt.Println("Slice",numbers)
	fmt.Println("Lenght",len(numbers))
	fmt.Println("Capacity",cap(numbers)) //when we append the element in slice and the length exceeds the capacity then the capacity is doubled
}