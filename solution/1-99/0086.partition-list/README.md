# [86. 分隔链表](https://leetcode-cn.com/problems/partition-list)

### 题目描述

<p>给定一个链表和一个特定值<em> x</em>，对链表进行分隔，使得所有小于 <em>x</em> 的节点都在大于或等于 <em>x</em> 的节点之前。</p>

<p>你应当保留两个分区中每个节点的初始相对位置。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> head = 1-&gt;4-&gt;3-&gt;2-&gt;5-&gt;2, <em>x</em> = 3
<strong>输出:</strong> 1-&gt;2-&gt;2-&gt;4-&gt;3-&gt;5
</pre>



### 解题思路

1. 双指针

### 具体解法

#### **Golang**
```go
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
```

