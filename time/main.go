package main

import (
	"fmt"
	"time"
)
func main(){
	current_time :=time.Now()
	formatted:=current_time.Format("2006-01-02,Monday")
	println("current time is",formatted) 
	layout := "2006-01-02"
	dateStr:= "2021-01-01"
	formatted_time,_ :=time.Parse(layout,dateStr)
	fmt.Println("formatted time is",formatted_time)


	new_date :=current_time.Add(24*time.Hour)
	fmt.Println("new date is",new_date)
	formatted_new_date :=new_date.Format("2006-01-02,Monday")
	fmt.Println("formatted new date is",formatted_new_date)
}