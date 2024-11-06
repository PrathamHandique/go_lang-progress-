package main

import "fmt"
func main(){
	day := 4
	switch day{
	case 1,2,4:
		println("Monday")
	case 3:
		println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	temp := 32
	switch {
	case temp < 0:
		fmt.Println("Freezing")
	case temp > 0 && temp < 20:
		fmt.Println("Cold")
	case temp > 20 && temp < 30:
		fmt.Println("Normal")
	case temp > 30:
		fmt.Println("Hot")
	default:
		fmt.Println("Invalid temperature")
}
}