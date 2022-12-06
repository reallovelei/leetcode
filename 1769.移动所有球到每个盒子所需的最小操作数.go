/*
 * @lc app=leetcode.cn id=1769 lang=golang
 *
 * [1769] 移动所有球到每个盒子所需的最小操作数
 */

// @lc code=start
func minOperations(boxes string) []int {
	// 先计算出
	left, right, cnt := int(boxes[0]-'0'), 0, 0
	for i := 1; i < len(boxes); i++ {
		if boxes[i] == '1' {
			right++
			cnt += i
		}
	}

	// fmt.Printf("left:%d right:%d cnt:%d\n", left, right, cnt)
	rs := make([]int, len(boxes))
	rs[0] = cnt
	for i := 1; i < len(boxes); i++ {
		rs[i] = rs[i-1] + left - right
		// fmt.Printf("left:%d right:%d rs:%+v\n", left, right, rs)
		if boxes[i] == '1' {
			right--
			left++
		}
	}
	return rs
}

// @lc code=end

