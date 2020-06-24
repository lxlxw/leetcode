package leetcode

import (
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {

	listNode1 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	listNode2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 8,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	getIntersectionNode(listNode1, listNode2)
}
