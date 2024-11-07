package main

import (
	"fmt"
	"net/url"
)
func main(){
	fmt.Println("handling urls")
	myUrl:="https://www.google.com"
	fmt.Println(myUrl)
	parsedUrl,err :=url.Parse(myUrl)
	if err!=nil{
		fmt.Println("error while parsing url",err)
		return
	}
	fmt.Println("scheme is",parsedUrl.Scheme)
	fmt.Println("host is",parsedUrl.Host)
	fmt.Println("path is",parsedUrl.Path)
	fmt.Println("query is",parsedUrl.RawQuery)
	//modifying url components
	parsedUrl.Scheme="http"
	parsedUrl.Host="localhost:8080"
	parsedUrl.Path="/search"
	parsedUrl.RawQuery="q=golang"
	newUrl :=parsedUrl.String()
	fmt.Println("new url is",newUrl)

}