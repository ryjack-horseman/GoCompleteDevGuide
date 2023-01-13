package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

type contactinfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactinfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("John")
	jim.print()
}

func (ptr2Person *person) updateName(newFirstName string) {
	(*ptr2Person).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
