package leetcode

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

// @lc code=end
