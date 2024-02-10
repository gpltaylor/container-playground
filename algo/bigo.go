package main

func Sum_Char_Codes(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c) - 48
	}
	return sum
}

// Create a bubble sort function - Big O(n^2)
func Bubble_Sort(arr []int) ([]int, int) {
	count := 0

	// create a variable to track if we made a swap
	swapped := true
	// loop through the array until we don't make a swap
	for swapped {
		// set the swap variable to false
		swapped = false
		// loop through the array
		for i := 0; i < len(arr)-1; i++ {
			count++
			// if the current element is greater than the next
			if arr[i] > arr[i+1] {
				// swap the elements
				arr[i], arr[i+1] = arr[i+1], arr[i]
				// set the swap variable to true
				swapped = true
			}
		}
	}
	// return the sorted array
	return arr, count
}
