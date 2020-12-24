package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int8
	contact   contactInfo
}

type contactInfo struct {
	email  string
	mobile int
}

func (p person) print() {
	fmt.Printf(
		"Name: %s %s\nAge: %d\n✉: %s | ☎: %d\n\n",
		p.firstName, p.lastName,
		p.age,
		p.contact.email,
		p.contact.mobile,
	)
}

// Note: Go is Pass By Value language, so below function will not
// work as expected
func (p person) updateLastName(newName string) {
	p.lastName = newName
}

// Pointer use case
func (p *person) updateAge(newAge int8) {
	(*p).age = newAge
}

func main() {
	bob := person{"Bob", "Sailor", 30, contactInfo{"abc@abc.cc", 1234567890}} // what if we change the struct defination, so do not use this syntax
	bob.print()

	ryan := person{
		firstName: "Ryan",
		lastName:  "Pop",
		age:       35,
		contact:   contactInfo{"pqr@pqr.dd", 9090909090},
	}
	ryan.print()

	var dave person
	dave.print() // It should print Zero values for all fields
	dave.firstName = "Dave"
	dave.lastName = "Zero"
	dave.age = 40
	dave.contact.email = "dave@yahoo.com"
	dave.contact.mobile = 9898098980
	dave.print() // Displays kind of key value

	// Unexpected behaviour as it does not update the last name
	// in the current person type variable
	dave.updateLastName("John")
	dave.print()

	// Use pointer to update the values as above is not working
	personPinter := &dave
	personPinter.updateAge(75)
	dave.print()

	// ShortCut, GoLang will automatically consider reference type to pointer of type
	dave.updateAge(80)
	dave.print()

}
