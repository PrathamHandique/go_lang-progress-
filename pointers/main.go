package main
import "fmt"



func modifyValueByReference(num *int){
	*num = *num + 2
}	


func main(){
	fmt.Println("we will learn about pointers")
	num :=10
	ptr := &num
	fmt.Println("pter contains",ptr)
	fmt.Println("value at ptr is",*ptr)
	value :=11
	modifyValueByReference(&value)
	fmt.Println("value after modification is",value)
}