package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

// func main() {
// 	alex := person{firstName: "Alex", lastName: "Anderson"}
// 	fmt.Println(alex)
// }

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "party",
		contactInfo: contactInfo{
			email:   "Jim@gmail.com",
			zipcode: 94000,
		},
	}
	//jimPointer := &jim
	jim.updateName("Jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

//	func (p *person) updateName(newFirstName string) {
//		p.firstName = newFirstName
//	}
func (p person) print() {
	fmt.Printf("%+v", p)
}
