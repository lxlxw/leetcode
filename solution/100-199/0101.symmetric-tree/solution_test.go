package leetcode

import (
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	var ret bool
	p := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}

	ret = false
	if ret != isSymmetric(p) {
		t.Fatalf("case fails %v\n", ret)
	}

	p1 := &TreeNode{}
	ret = true
	if ret != isSymmetric(p1) {
		t.Fatalf("case fails %v\n", ret)
	}

}
