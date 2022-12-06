/*
 * @lc app=leetcode.cn id=1758 lang=golang
 *
 * [1758] 生成交替二进制字符串的最少操作数
 */

// @lc code=start
func minOperations(s string) int {
	cnt := 0
	for k, v := range s {
		if k%2 != int(v-'0') {
			cnt++
		}
	}
	return min(cnt, len(s)-cnt)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end

