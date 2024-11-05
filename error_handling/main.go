package main

import "fmt"
//method1
// func divide(a,b float64)(float64){
// 	 if b==0{
// 		return 0
// 	 }
// 	 return a/b 
// }

//method2
func divide(a,b float64)(float64, error){
	if b==0{
	   return 0, fmt.Errorf("cannot divide by zero")
	}
	return a/b , nil
}
func main(){
	fmt.Println("learning error handling")
	//ans:= divide(10,0) //valid for method1
	//ans, _ := divide(10,0) //helps to ignore the error
	ans , err := divide(10,3) 
	if err != nil{
		fmt.Println("handling error")
	}
	fmt.Printf("The division is %.3f\n", ans) 
	fmt.Println("The divisoin is ", ans) //this demonstrates the difference between Printf and Println
}