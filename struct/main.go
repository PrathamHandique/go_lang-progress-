package main

import "fmt"

type Person struct{
	FirstName string
	LastName string
	Age int
}
type Contact struct{
	Email string
	Phone string
}
type Address struct{
	House int
	Area string
	State string
}
type Employees struct{
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}
func main(){
	fmt.Println("we will learn about structs")
	var prince Person //creating a variable of type Person
	fmt.Println(prince)// we will get the result like { 0} because it is space space and a zero for integer
	prince.FirstName = "Prince"
	prince.LastName = "Kumar"
	prince.Age = 21
	fmt.Println(prince) 
	//direct assigning values to struct
	person1 := Person{
		FirstName: "harry",
		LastName: "Kumar",
		Age: 21,
	}
	fmt.Println(person1)
	//new method
	person2 := new(Person)
	person2.FirstName = "harry"
	person2.LastName = "sejal"
	person2.Age = 21
	fmt.Println(person2)
	fmt.Println(person2.Age)

	//nested struct
	var employee1 Employees
	employee1.Person_Details =Person{
		FirstName: "harry",
		LastName: "Kumar",
		Age: 21,
	}

	employee1.Person_Contact.Email = "hello@gmail.com"
	employee1.Person_Contact.Phone = "1234567890"
	fmt.Println(employee1)
	fmt.Println(employee1.Person_Contact.Email)

}