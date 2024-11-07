package main

import (
	"encoding/json"
	"fmt"
)
//method to convert struct into json
type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
	IsAdult bool `json:"is_adult"`
}
func main(){
	person :=Person{Name:"Prince",Age:21,IsAdult:true}
	fmt.Println(person)
	//convert person into json uncoding(marshal)
	jsonData,err :=json.Marshal(person)
	if err!=nil{
		fmt.Println("error while encoding json",err)
		return
	}
	fmt.Println("the person data in json is",string(jsonData))
	//decoding json data(unmarshal)
	var personData Person 
	err =json.Unmarshal(jsonData,&personData)
	if err!=nil{
		fmt.Println("error while decoding json",err)
		return
	}
	fmt.Println("the person data is",personData)


}