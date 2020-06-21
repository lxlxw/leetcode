# [83.删除排序链表中的重复元素](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)


### 题目描述

<div class="notranslate"><p>给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:</strong> 1-&gt;1-&gt;2
<strong>输出:</strong> 1-&gt;2
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> 1-&gt;1-&gt;2-&gt;3-&gt;3
<strong>输出:</strong> 1-&gt;2-&gt;3</pre>
</div>

### 解题思路

1. 当前节点的下一个指针不为空时，判断当前节点的值跟下一个指针节点的值是否相等。
2. 如果相等，则将下下个节点的指针指向当前节点的下一个指针。
3. 如果不相等，则指针往前移。

### 代码实现

<!-- tabs:start -->

#### **Golang**
```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ret := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}
	return ret
}
```


<!-- tabs:end -->