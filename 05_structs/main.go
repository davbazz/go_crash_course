package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

// value reciever
func (p Person) greet() string {
	return "Greetings to " + p.firstName + " " + p.lastName + " " + "who is " + strconv.Itoa(p.age) + " " + "old"
}

// pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

// pointer receiver
func (p *Person) getMarried() {
	p.lastName = "bazashvili"
}

func main() {
	user := Person{"anna", "smiyan", 23}
	fmt.Println(user.greet())

	// user.age = 23
	// fmt.Println(user.greet())

	user.hasBirthday()
	fmt.Println(user.greet())

	user.getMarried()
	fmt.Println(user.greet())
}
