/*
 * @lc app=leetcode.cn id=1796 lang=golang
 *
 * [1796] 字符串中第二大的数字
 */
// import "unicode"

// @lc code=start
func secondHighest(s string) int {
	max1, max2 := -1, -1

	for _, v := range s {
		if unicode.IsDigit(v) {
			val := int(v - '0')
			if val > max1 {
				max1, max2 = val, max1
			} else if max1 > val && val > max2 {
				max2 = val
			}
		}
	}
	return max2
}

// @lc code=end

