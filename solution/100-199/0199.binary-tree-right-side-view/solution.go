package leetcode

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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

// @lc code=end
