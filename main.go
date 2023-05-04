// package main

// import "fmt"

// type User struct {
// 	firstName, lastName, location string
// 	age                           int
// }

// func (u *User) printUser() {
// 	fmt.Printf("Hi, my name is %v %v, I'm %v years old and I live in %v", u.firstName, u.lastName, u.age, u.location)
// }

// func main() {
// 	u := &User{"david", "bazashvili", "london", 22}
// 	fmt.Println(u)

// 	u.printUser()
// }

package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName, lastName, location string
	age                           int
}

func (u *User) changeLastName(lastName string) {
	u.lastName = lastName
}

// func (u *User) changeAge(date string) {
// 	if time.Date() == "22.09.2000" {
// 		&u.age++
// 	}
// }

func main() {
	u := &User{"david", "bazashvili", "london", 22}
	fmt.Println(u)

	u.changeLastName("new")
	fmt.Println(u)

	fmt.Println(time.Now())
}
