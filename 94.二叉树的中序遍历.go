/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
//func inorderTraversal(root *TreeNode) (rs []int) {
//	stack := []*TreeNode{}
//	if root == nil {
//		return
//	}
//
//	cur := root
//	for cur != nil || len(stack) > 0 {
//		for cur != nil {
//			stack = append(stack, cur)
//			cur = cur.Left
//		}
//		// 拿最后一个 类似 pop
//		cur = stack[len(stack)-1]
//		// 把stack 最后去掉。
//		stack = stack[:len(stack)-1]
//		rs = append(rs, cur.Val)
//		cur = cur.Right
//	}
//	return rs
//}
type ColorNode struct {
	Node  *TreeNode
	Color int
}

func inorderTraversal(root *TreeNode) (rs []int) {
	// 对根节点 进行判空
	if root == nil {
		return
	}
	stack := []ColorNode{Node: root, Color: 0}
	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 如果已经着色 那么进入结果, 每个节点第一次进栈是不会进结果的。
		if last.Color == 1 {
			rs = append(rs, last.Node.Val)
		} else {
			if last.Node.Right != nil {
				stack = append(stack, last.Node.Right)
			}

			stack = append(stack, ColorNode{Node: root, Color: 1})

			if last.Node.Left != nil {
				stack = append(stack, last.Node.Left)
			}
		}
	}
	return rs
}

// @lc code=end

