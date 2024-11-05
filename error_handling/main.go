package main

import "fmt"

func divide(a,b float64)(float64){
	 if b==0{
		return 0
	 }
	 return a/b 
}
func main(){
	fmt.Println("learning error handling")
	ans:= divide(10,0)
	fmt.Println("The division is ", ans)
	
}