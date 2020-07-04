package leetcode

import (
	"testing"
)

func TestRotate(t *testing.T) {
	var matrix [][]int

	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
}
