package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"encoding/json"
)

type Todo struct{
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func main(){
	fmt.Println("Hello World")
	res, err :=http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode !=http.StatusOK{
		fmt.Println("Error: status code", res.StatusCode)
		return
	}
	
	data, err:= ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(data))
	
	
	// var todo Todo
	// err =json.NewDecoder(res.Body).Decode(&todo)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// 	fmt.Println(todo)
}