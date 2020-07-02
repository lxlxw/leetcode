# [101. 对称二叉树](https://leetcode-cn.com/problems/symmetric-tree/description/)

### 题目描述

<p>给定一个二叉树，检查它是否是镜像对称的。</p>

<p>&nbsp;</p>

<p>例如，二叉树&nbsp;<code>[1,2,2,3,4,4,3]</code> 是对称的。</p>

<pre>    1
   / \
  2   2
 / \ / \
3  4 4  3
</pre>

<p>&nbsp;</p>

<p>但是下面这个&nbsp;<code>[1,2,2,null,3,null,3]</code> 则不是镜像对称的:</p>

<pre>    1
   / \
  2   2
   \   \
   3    3
</pre>

<p>&nbsp;</p>

<p><strong>进阶：</strong></p>

<p>你可以运用递归和迭代两种方法解决这个问题吗？</p>

### 解题思路

1. 递归解法

### 具体解法


#### **Golang**
```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recursive(root.Left, root.Right)
}
func recursive(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	return p.Val == q.Val && recursive(p.Left, q.Right) && recursive(p.Right, q.Left)
}
```


