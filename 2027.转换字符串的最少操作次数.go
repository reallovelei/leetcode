/*
 * @lc app=leetcode.cn id=2027 lang=golang
 *
 * [2027] 转换字符串的最少操作次数
 */

// @lc code=start
func minimumMoves(s string) int {
	pos, max, n := 0, 0, len(s)
	// 从头遍历，遇到X 就往后挪3位
	for pos < n {
		if s[pos] == 'X' {
			max++
			pos += 3
			continue
		}
		pos++

	}
	return max
}

// @lc code=end

