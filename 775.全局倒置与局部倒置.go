/*
 * @lc app=leetcode.cn id=775 lang=golang
 *
 * [775] 全局倒置与局部倒置
 */

// @lc code=start
func isIdealPermutation(nums []int) bool {

	for k, v := range nums {
		if abs(v-k) > 1 { // 这里是 > 1
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

