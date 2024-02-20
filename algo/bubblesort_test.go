package main

import (
	"fmt"
	"testing"
)

func BubbleSort(arr []int) []int {
	swapped := false
	ln := len(arr) - 1

	for {
		swapped = false
		for i := 0; i < ln; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		ln--

		if !swapped {
			break
		}
	}
	return arr
}

func BubbleSortPrimagen(arr []int) []int {
	ln := len(arr)

	for i := 0; i < ln; i++ {
		for j := 0; j < ln-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			fmt.Println(arr, i, j, arr[j] < arr[j+1])
		}
	}

	return arr
}

func Test_BubbleSort(t *testing.T) {
	arr := []int{4, 5, 3, 9, 8}
	sorted := BubbleSortPrimagen(arr)
	if sorted[0] != 3 {
		t.Error("Expected 333got ", sorted[0])
	}
	if sorted[1] != 4 {
		t.Error("Expected 4, got ", sorted[1])
	}
	if sorted[2] != 5 {
		t.Error("Expected 5, got ", sorted[2])
	}
	if sorted[3] != 8 {
		t.Error("Expected 8, got ", sorted[3])
	}
	if sorted[4] != 9 {
		t.Error("Expected 9, got ", sorted[4])
	}
}

func Benchmark_BubbleSort_gpltaylor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort([]int{5, 4, 3, 2, 1})
	}
}

func Benchmark_BubbleSort_Primagen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSortPrimagen([]int{5, 4, 3, 2, 1})
	}
}
