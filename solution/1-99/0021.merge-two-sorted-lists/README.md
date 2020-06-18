# [21.合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists)


### 题目描述

<div class="notranslate"><p>将两个升序链表合并为一个新的 <strong>升序</strong> 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。&nbsp;</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre><strong>输入：</strong>1-&gt;2-&gt;4, 1-&gt;3-&gt;4
<strong>输出：</strong>1-&gt;1-&gt;2-&gt;3-&gt;4-&gt;4
</pre>
</div>

### 解题思路
![](http://lc-photo.xwlin.com/21-1.png)
![](http://lc-photo.xwlin.com/21-2.png)
![](http://lc-photo.xwlin.com/21-3.png)

### 代码实现

<!-- tabs:start -->

#### **Golang 1**
```go
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
```

#### **Golang 2**
```go
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
```

<!-- tabs:end -->