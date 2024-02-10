package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSort(t *testing.T) {
	ans, count := Bubble_Sort([]int{5, 3, 8, 6, 2, 9, 4, 5, 6, 8})

	require.Equal(t, 45, count, "The count should be 10.")
	require.Equal(t, ans, []int{2, 3, 4, 5, 5, 6, 6, 8, 8, 9}, "The list should be sorted.")
}
