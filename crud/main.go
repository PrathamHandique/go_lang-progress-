package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	//"encoding/json"
)

type Todo struct{
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func performGetRequest(){
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

func performPostRequest(){
	todo:=Todo{
		UserId:23,
		Id:1,
		Title:"Hello World",
		Completed:false,
	}
	//convert todo into json
	jsonData,err :=json.Marshal(todo) 
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//convert json data to string
	jsonString :=string(jsonData)
	//convert json  data to reader
	jsonReader :=strings.NewReader(jsonString)
	myUrl := "https://jsonplaceholder.typicode.com/todos"
	//send post request
	res, err :=http.Post(myUrl,"application/json",jsonReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()
	data, _:= ioutil.ReadAll(res.Body)
	fmt.Println(string(data))


}

func performUpdateRequest(){
	todo :=Todo{
		UserId:23224,
		Id:1,
		Title:"Hello Pratham",
		Completed:true,
	}
	jsonData,err :=json.Marshal(todo) 
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonString :=string(jsonData)
	//convert json  data to reader
	jsonReader :=strings.NewReader(jsonString)
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"
	//create put request
	req, err :=http.NewRequest(http.MethodPut,myUrl,jsonReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req.Header.Set("Content-Type","application/json")
	client :=http.Client{}
	res, err :=client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()
	data ,_:= ioutil.ReadAll(res.Body)
	fmt.Println(string(data))


	
}

func performDeleteRequest(){
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"
	//create delete request
	req, err :=http.NewRequest(http.MethodDelete,myUrl,nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	client :=http.Client{}
	res, err :=client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Deleted successfully status",res.StatusCode)
}

func main(){
	fmt.Println("Hello World")
	performGetRequest()
	performPostRequest()
	performUpdateRequest()
	performDeleteRequest()

	
}