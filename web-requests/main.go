package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main(){
	res,err :=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err!=nil{
		fmt.Println("error while fetching data",err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err !=nil{
		fmt.Println("error while reading data",err)
		return
	}
	fmt.Println(string(data)) // we need to write string(data) because data is of type []byte 
}