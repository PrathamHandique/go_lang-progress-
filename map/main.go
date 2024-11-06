package main

import "fmt"
func main(){
	fmt.Println("learning maps")
	studentGrades := make(map[string]int)
	studentGrades["Timmy"] = 42
	studentGrades["Jess"] = 92
	studentGrades["Sam"] = 67
	fmt.Println("marks of timmy is ",studentGrades["Timmy"])
	delete(studentGrades,"Timmy")
	fmt.Println("marks of timmy is ",studentGrades["Timmy"])
	//checking if a key exists
	//if a key does not exist, the value returned is 0
	grades,exists := studentGrades["Jess"]
	fmt.Println("grade of jess is ",grades)
	fmt.Println("does jess exist ",exists)

	person :=map[string]int{
		"tammy": 42,
		"jess": 92,
	}
	for index,value :=range person{
		fmt.Println(index,value)
	}
	}
