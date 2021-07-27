package main

import (
	"fmt"
	"strconv"
)

// Person Define Person Struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	//Alternative
	// firstName, lastName, city, gender, age string
	//age int
}

// Greeting method - value receiver --- doesnt change anything
func (person Person) greet() string {
	return "Hello. my name is " + person.firstName + " " + person.lastName +
		" and I am " + strconv.Itoa(person.age)
}

//Has birthday - pointer receiver --- change something
func (person *Person) hasBirthday() {
	person.age++
}

//get Married - pointer receiver
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "M" {
		return

	} else {
		person.lastName = spouseLastName
	}

}

func main() {
	person1 := Person{firstName: "Joel", lastName: "Kanyi", city: "Narok", gender: "M", age: 23}

	person2 := Person{firstName: "Brandy", lastName: "Odhiambo", city: "Kisumu", gender: "F", age: 22}

	//Alternative
	//person2 := Person{ "Joel",  "Kanyi",  "Narok", "M",  23}

	//fmt.Println(person1)
	//
	//fmt.Println(person1.firstName)
	//
	//person1.age++
	//fmt.Println(person1.age)
	//person1.hasBirthday()
	fmt.Println(person1.greet())
	person2.getMarried("Kanyi")
	fmt.Println(person2.greet())
}
