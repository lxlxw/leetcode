# [199. 二叉树的右视图](https://leetcode-cn.com/problems/binary-tree-right-side-view/description/)

### 题目描述

<p>给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。</p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong>&nbsp;[1,2,3,null,5,null,4]
<strong>输出:</strong>&nbsp;[1, 3, 4]
<strong>解释:
</strong>
   1            &lt;---
 /   \
2     3         &lt;---
 \     \
  5     4       &lt;---
</pre>

### 解题思路

1. BFS

### 具体解法


#### **Golang**
```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	res = []int{root.Val}
	for len(queue) > 0 {
		l := len(queue)
		flag := 0
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Right != nil {
				if flag == 0 {
					flag = node.Right.Val
				}
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				if flag == 0 {
					flag = node.Left.Val
				}
				queue = append(queue, node.Left)
			}
		}
		if flag != 0 {
			res = append(res, flag)
		}
		queue = queue[l:]
	}
	return res
}
```


