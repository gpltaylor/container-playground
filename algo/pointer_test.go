package main

import "testing"

func increment(n int) {
	n++
}

func incrementPointer(n *int) {
	*n++
}

func Test_Pointer(t *testing.T) {
	// create a pointer to an int
	var x int = 10
	var n *int
	n = &x

	increment(*n)

	// assert that n is still 10
	if *n != 10 {
		t.Error("Expected 10, got ", n)
	}

	incrementPointer(n)

	// assert that n is now 11
	if *n != 11 {
		t.Error("Expected 11, got ", n)
	}

}
