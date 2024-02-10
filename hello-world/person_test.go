package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerson(t *testing.T) {
	// create a new Person Struct
	p := Person{
		FirstName: "Jane",
	}

	require.Equal(t, p.FirstName, "Jane", "First name should be set in constructor")

	// set the first and last name
	p.SetFirstName("John")
	p.SetLastName("Doe")

	require.Equal(t, p.FirstName, "John", "First name should be John")
	require.Equal(t, p.LastName, "Doe", "Last name should be {Doe}")

	// Test Getters
	require.Equal(t, p.GetFirstName(), "John", "The two words should be the same.")
	require.Equal(t, p.GetLastName(), "Doe", "The two words should be the same.")

}
