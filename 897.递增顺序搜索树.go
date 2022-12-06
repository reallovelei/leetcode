/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序搜索树
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
func increasingBST(root *TreeNode) *TreeNode {
	var arr []int
	var midFunc func(*TreeNode)
	midFunc = func(node *TreeNode) {
		// 按中序遍历的形式 将数放到数组中
		if node != nil {
			midFunc(node.Left)
			arr = append(arr, node.Val)
			midFunc(node.Right)
		}
	}
	midFunc(root)

	r := &TreeNode{}
	curNode := r
	// 开始将数组中的数 挂到目标树上。
	for _, v := range arr {
		curNode.Right = &TreeNode{Val: v}
		curNode = curNode.Right
	}
	return r.Right
}

// @lc code=end

