# [111. 二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/)

### 题目描述

<p>给定一个二叉树，找出其最小深度。</p>

<p>最小深度是从根节点到最近叶子节点的最短路径上的节点数量。</p>

<p><strong>说明:</strong>&nbsp;叶子节点是指没有子节点的节点。</p>

<p><strong>示例:</strong></p>

<p>给定二叉树&nbsp;<code>[3,9,20,null,null,15,7]</code>,</p>

<pre>    3
   / \
  9  20
    /  \
   15   7</pre>

<p>返回它的最小深度 &nbsp;2.</p>

### 解题思路


### 具体解法

1. BFS

#### **Golang**
```go
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				return count
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		count++
		queue = queue[l:]
	}
	return count
}
```


