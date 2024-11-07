package main

import (
	"fmt"
	"os"
)
func main(){
	/*
	//file,err :=os.Create("file-handling/test.txt")
	file,err :=os.Create("file-handling/test1.txt")
	if err !=nil{
		fmt.Println("error creating file",err)
	}
	defer file.Close()
	content :="hello world"
	//_,err =io.WriteString(file,content + "\n")  //\n is used to write to a new line in this case to put the curser to the next line
	byte,err :=io.WriteString(file,content + "\n")
	fmt.Println("bytes written are",byte)
	if err !=nil{
		fmt.Println("error writing to file",err)
		return
	}
	fmt.Println("file written successfully")
	*/

	//reading from the file
	/*
	file ,err:=os.Open("file-handling/test1.txt")
	if err !=nil{
		fmt.Println("error opening file",err)
		return
	}
	defer file.Close()
	//create a buffer to read the file
	buffer :=make([]byte,1024)
	for{
		n,err :=file.Read(buffer)
		if err == io.EOF{
			break
		}
		if err !=nil{
			fmt.Println("error reading file",err)
			return
		}
		fmt.Println(string(buffer[:n]))  //[:n] is used to print only the number of bytes read because n is the number of bytes read and we use the :n method to pr
		//another mehtod to print the file
		//fmt.Println(string(buffer))
		fmt.Println(n)
	}
	*/
	//another method to read the file
	//file ,err:=ioutil.ReadFile("file-handling/test1.txt")// depricated method
	file,err :=os.ReadFile("file-handling/test1.txt") //new method
	if err !=nil{
		fmt.Println("error reading file",err)
		return
	}
	fmt.Println(string(file))



}