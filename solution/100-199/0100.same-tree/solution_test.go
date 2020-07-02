package leetcode

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestIsSameTree(t *testing.T) {
	var ret bool

	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	ret = true
	if ret != isSameTree(p, q) {
		t.Fatalf("case fails %v\n", ret)
	}

	p1 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	q1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	ret = false
	if ret != isSameTree(p1, q1) {
		t.Fatalf("case fails %v\n", ret)
	}

	p2 := &TreeNode{}
	q2 := &TreeNode{}
	ret = true
	if ret != isSameTree(p2, q2) {
		t.Fatalf("case fails %v\n", ret)
	}
}
