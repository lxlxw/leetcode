package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	listNode1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	listNode2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	ret := addTwoNumbers(listNode1, listNode2)
	if ret.Val != 7 || ret.Next.Val != 0 || ret.Next.Next.Val != 8 {
		t.Fatalf("case fails %v\n", ret)
	}
}
