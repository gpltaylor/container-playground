package main

import "testing"

func Test_BsSearch(t *testing.T) {
	type testCase struct {
		name     string
		elements []int
		target   int
		expected bool
	}

	tests := []testCase{
		{"Test with target at index 0 with duplicate", []int{1, 3, 5, 7, 9, 14, 16, 36, 79, 89, 90}, 89, true},
		{"Test with target at the beginning", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, true},
		{"Test with target at the end", []int{1, 2, 3}, 2, true},
	}

	for i := 0; i < len(tests); i++ {
		t.Run(tests[i].name, func(t *testing.T) {
			ans := binarySearch(tests[i].elements, tests[i].target)
			if ans != tests[i].expected {
				t.Errorf("Test failed, expected: '%t', got:  '%t'", tests[i].expected, ans)
			}
		})
	}

}
