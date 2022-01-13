package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	//shorter way below
	//contact   contactInfo
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		//shorter way below
		//contact: contactInfo{
		contactInfo: contactInfo{
			email:   "abc@123.gov",
			zipCode: 12345,
		},
	}

	//jimPointer := &jim
	//jimPointer.updateName("jimbo")
	jim.updateName("jimbo")
	jim.print()

}

/*
	Course Notes

	RECEIVER -- (<value, type>)
	(*person) -- an * in front of a type is a type description, we're working with a pointer
	to a person. this updateName function can only be called with a pointer to a person.

	(*pointerToPerson) -- this is an operator, it means we want to manipulate the value
	that the pointer is referencing. *pointerToPerson takes the pointer and turns it into
	the actual value.
*/
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
