package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArray(t *testing.T) {

	x := GetArray()

	require.Equal(t, len(x), 3, "The len on the array should be 3.")

	a := GetMatrix()
	require.Equal(t, len(a), 2, "The len on the matrix should be 2.")

}
