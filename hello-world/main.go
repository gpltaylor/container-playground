// create an hello world print statement in go

package main

import (
	"fmt"
)

func main() {
	// create a new Person Struct
	p := Person{
		FirstName: "Jane",
	}

	// set the first and last name
	p.SetFirstName("John")

	p.SetLastName("Doe")

	// print the first and last name
	fmt.Println(p.FirstName, p.LastName)
	PrintPerson(&p)

	fmt.Println("Hello World!")
}
