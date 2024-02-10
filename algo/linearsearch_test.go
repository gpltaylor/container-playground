package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LinearSearch(t *testing.T) {
	// Define a struct for the test cases
	type testCase struct {
		name     string
		elements []int
		target   int
		expected int
	}

	// Create a slice of test cases
	tests := []testCase{
		{"Test with target at index 0 with duplicate", []int{6, 3, 8, 6, 2, 9, 4, 5, 6, 8}, 6, 0},
		{"Test with target at the beginning", []int{6, 3, 8, 6, 2, 9, 4, 5, 6, 8}, 6, 0},
		{"Test with target at the end", []int{5, 3, 8, 1, 2, 9, 4, 5, 1, 6}, 6, 9},
		{"Test with no target", []int{5, 3, 8, 1, 2, 9, 4, 5, 7, 8}, 6, -1},
		// Add more test cases as needed
	}

	for i := 0; i < len(tests); i++ {
		// Run the test case
		t.Run(tests[i].name, func(t *testing.T) {
			ans := Linear_Search(tests[i].elements, tests[i].target)
			require.Equal(t, tests[i].expected, ans, tests[i].name)
		})
	}
}
