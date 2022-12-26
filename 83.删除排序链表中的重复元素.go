/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead *ListNode
	var newPrev *ListNode

	newHead.Val, newPrev.Val = head.Val, head.Val

	prev := head
	// head 后移一步 与 prev 形成相邻

	head = head.Next

	for head != nil {
		if head.Val == prev.Val {
			head = head.Next
			continue
		}

		tmp := prev
		newPrev.Next = prev

		prev = head
		head = head.Next
	}
	return newHead
}

// @lc code=end

