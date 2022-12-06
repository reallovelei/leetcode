/*
 * @lc app=leetcode.cn id=1624 lang=golang
 *
 * [1624] 两个相同字符之间的最长子字符串
 */

// @lc code=start
func maxLengthBetweenEqualCharacters(s string) int {
	rs := -1
	// 定义一个以字母为key index为val的数组.
	mapIndex := [26]int{}
	// 初始化这个数组的val都为-1
	for i := range mapIndex {
		mapIndex[i] = -1
	}

	for k, v := range s {
		d := v - 'a'
		// 第一次扫描到
		if mapIndex[d] < 0 {
			mapIndex[d] = k
		} else {
			// 比较一个较大
			rs = max(rs, k-mapIndex[d]-1)
		}
	}
	return rs
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

