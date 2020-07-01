package leetcode

import (
	"testing"
)

func TestPartition(t *testing.T) {
	var x int
	x = 3
	listNode := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  2,
							Next: nil,
						},
					},
				},
			},
		},
	}
	partition(listNode, x)
}
