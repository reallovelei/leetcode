/*
 * @lc app=leetcode.cn id=1750 lang=golang
 *
 * [1750] 删除字符串两端相同字符后的最短长度
 */

// @lc code=start
func minimumLength(s string) int {
	left, right := 0, len(s)-1
	// 判断第一次 是否进入。如果第一个字符 和最后一个都不一样就没有可以删除的
	for left < right && s[left] == s[right] {
		ch := s[left]
		//  拿到当前的 首字符 进行判断 且 前缀和后缀不能有重复
		// 注意：这个条件需要放到前面，否则 可能会出现数组越界的情况
		for left <= right && s[left] == ch {
			left++
		}

		for left <= right && s[right] == ch {
			right--
		}
	}

	return right + 1 - left
}

// @lc code=end

