package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I " + strconv.Itoa(p.age)
}

// pointer receiver (change something on a struct)
func (p *Person) hasBirthday() { // not returning anything so you don't have to specify
	p.age++
}

// another pointer receiver
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Sam", lastName: "Williams", city: "Boston", gender: "m", age: 20}

	// Alternative
	person2 := Person{"Jan", "Smith", "Boston", "f", 20}
	fmt.Println(person1)
	fmt.Println(person2)

	// fmt.Println(person1.firstName)
	person1.age = 22
	// fmt.Println(person1)

	//  --- 2 kinds of methods --- //
	// value receivers just do calcuations / don't change a value
	fmt.Println(person2.greet())

	// pointer receivers / for changing something
	person1.hasBirthday()
	person2.getMarried("Williams")
	fmt.Println(person2.greet())

}
