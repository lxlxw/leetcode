package leetcode

import (
	"testing"
)

func TestRightSideView(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	ret := []int{1, 2}
	nums := rightSideView(root)
	for i, v := range nums {
		if ret[i] != v {
			t.Fatalf("case fails %v\n", ret)
		}
	}

	root1 := &TreeNode{}
	nums1 := rightSideView(root1)
	if len(nums1) > 1 {
		t.Fatalf("case fails %v\n", nums1)
	}

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}
	ret2 := []int{1, 3, 6}
	nums2 := rightSideView(root2)
	for i, v := range nums2 {
		if ret2[i] != v {
			t.Fatalf("case fails %v\n", nums2)
		}
	}

}
