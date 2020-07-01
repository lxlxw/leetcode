package leetcode

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	lessTop := &ListNode{}
	less := lessTop
	greaterTop := &ListNode{}
	greater := greaterTop
	for n := head; n != nil; n = n.Next {
		if n.Val < x {
			less.Next = n
			less = n
			continue
		}
		greater.Next = n
		greater = n
	}
	less.Next = greaterTop.Next
	greater.Next = nil
	return lessTop.Next
}

// @lc code=end
