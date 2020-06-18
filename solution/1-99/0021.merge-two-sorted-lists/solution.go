package leetcode

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 */

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	ret := &ListNode{}
	if l1.Val <= l2.Val {
		ret = l1
		ret.Next = mergeTwoLists1(l1.Next, l2)
	} else {
		ret = l2
		ret.Next = mergeTwoLists1(l1, l2.Next)
	}
	return ret
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	ret := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			ret.Next = l1
			l1 = l1.Next
		} else {
			ret.Next = l2
			l2 = l2.Next
		}
		ret = ret.Next
	}
	if l1 != nil {
		ret.Next = l1
	}
	if l2 != nil {
		ret.Next = l2
	}
	return head.Next
}

// @lc code=end
