package leetcode

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	var ret int
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

	ret = 3
	if ret != maxDepth(p) {
		t.Fatalf("case fails %v\n", ret)
	}

}
