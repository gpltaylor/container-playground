package main

import "fmt"

// Create a Struct for a Person with First and Last Name
type Person struct {
	FirstName string
	LastName  string
}

// create getters and setters for the Person Struct
func (p *Person) SetFirstName(firstName string) {
	p.FirstName = firstName
}

func (p *Person) SetLastName(lastName string) {
	p.LastName = lastName
}

// create Getters for the Person Struct
func (p *Person) GetFirstName() string {
	return p.FirstName
}

func (p *Person) GetLastName() string {
	return p.LastName
}

// Create an interface for the Person Struct
type PersonInterface interface {
	SetFirstName(firstName string)
	SetLastName(lastName string)
	// Add Getters
	GetFirstName() string
	GetLastName() string
}

// Create a function that takes in a PersonInterface and print the first and last name
func PrintPerson(p PersonInterface) {
	fmt.Println(p.GetFirstName(), p.GetLastName())
}
