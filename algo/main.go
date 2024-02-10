package main

// create a public main function
func main() {
	// print hello world
	// msg := "134"
	// println(msg, Sum_Char_Codes(msg))

	answer, _ := Bubble_Sort([]int{5, 3, 8, 2, 1, 4})

	// loop through the array and print each element
	for _, v := range answer {
		print(v)
	}

}
