package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("we will learn about strings")
	data:="apple.orage.banana.banana"
	count :=strings.Count(data,"banana")
	data1:="   hello world    "
	trimmed :=strings.TrimSpace(data1)
	str1 := "pratham"
	str2 := "handique"
	res :=strings.Join([]string{str1,str2}," ")
	fmt.Println(strings.Split(data,"."))
	fmt.Println(count)
	fmt.Println(trimmed)
	fmt.Println(res)
}