/*
 * @lc app=leetcode.cn id=1784 lang=golang
 *
 * [1784] 检查二进制字符串字段
 */

// @lc code=start
func checkOnesSegment(s string) bool {
	rs := true
	// 如果 1的前面出现了0 则代表 出现了1个以上 不连续的1.
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == 49 && s[i-1] == 48 {
			return false
		}
	}

	return rs
}

// @lc code=end

