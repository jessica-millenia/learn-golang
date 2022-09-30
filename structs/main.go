package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	person1 := person{
		firstName: "Hello",
		lastName:  "World",
		contactInfo: contactInfo{
			email:   "Hello@gmail.com",
			zipCode: 90,
		},
	}

	person1.print()
	person1.updateName("Hi", "Golang")
	person1.updateContactInfo("Hi@gmail.com", 100)
	fmt.Println("\nAfter Updating Data:")
	person1.print()
}

func (pointerToPerson *person) updateName(newFirstName string, newLastName string) {
	(*pointerToPerson).firstName = newFirstName
	(*pointerToPerson).lastName = newLastName
}

func (pointerToPerson *person) updateContactInfo(newEmail string, newZipCode int) {
	(*pointerToPerson).contactInfo.email = newEmail
	(*pointerToPerson).contactInfo.zipCode = newZipCode
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
