package main

import (
	"math"
	"testing"
)

func BinarySearch(arr []bool, needle bool) int {
	sq := math.Sqrt(float64(len(arr)))
	bs := int(sq)

	min := 0
	max := len(arr) / bs

	for x := min; x < max; x++ {
		if arr[max] == needle {
			for i := 0; i < bs; i++ {
				if arr[min+i] == needle {
					return min + i
				}
			}
		}

		min, max = max, max+bs
		if max > len(arr) {
			max = len(arr) - 1
		}
	}

	return -1
}

func BinarySearchPrimagen(breaks []bool, needle bool) int {
	lnth := int(len(breaks))
	jmpAmount := int(math.Sqrt(float64(lnth)))

	i := int(jmpAmount)

	for ; i < lnth; i += jmpAmount {
		if breaks[i] {
			break
		}
	}

	i -= int(jmpAmount)

	for j := 0; j < jmpAmount && i < lnth; {

		if breaks[i] {
			return i
		}

		j++
		i++
	}

	return -1
}

func Test_BlockSearch(t *testing.T) {
	type testCase struct {
		name     string
		elements []bool
		target   bool
		expected int
	}

	tests := []testCase{
		{"Test 1", []bool{false, false, false, false, false, true, true, true, true, true}, true, 5},
		{"Test 2", []bool{false, false, false, false, false, true, true, true, true, true}, true, 5},
		{"Test 3", []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true}, true, 15},
		{"Test 4", []bool{false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true}, true, 32},
		{"Test 5", []bool{false, false, false, false, false, false, false, true, true, true}, true, 7},
	}

	for i := 0; i < len(tests); i++ {
		t.Run(tests[i].name, func(t *testing.T) {
			//ans := BinarySearch(tests[i].elements, tests[i].target)
			ans := BinarySearchPrimagen(tests[i].elements, tests[i].target)

			if ans != tests[i].expected {
				t.Errorf("Test failed, expected: '%d', got:  '%d'", tests[i].expected, ans)
			}
		})
	}

}

// go test -bench=BinarySearch*
func Benchmark_BinarySearch_gpltaylor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch([]bool{false, false, false, false, false, true, true, true, true, true}, true)
	}
}

func Benchmark_BinarySearch_Primagen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearchPrimagen([]bool{false, false, false, false, false, true, true, true, true, true}, true)
	}
}
