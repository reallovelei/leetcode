/*
 * @lc app=leetcode.cn id=1732 lang=golang
 *
 * [1732] 找到最高海拔
 */

// @lc code=start
func largestAltitude(gain []int) int {
	last, max := 0, 0
	for _, v := range gain {
		last += v
		if last > max {
			max = last
		}
	}
	return max
}

// @lc code=end

