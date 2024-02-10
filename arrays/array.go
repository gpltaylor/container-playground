package main

import (
	"fmt"
)

func GetArray() []int {
	// b := [5]int{1, 2, 3, 4, 5}
	return []int{1, 2, 3}
}

func GetMatrix() [2][3]int {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	return twoD
}

func main() {
	fmt.Println("Hello, playground")
}
