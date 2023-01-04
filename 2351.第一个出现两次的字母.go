/*
 * @lc app=leetcode.cn id=2351 lang=golang
 *
 * [2351] 第一个出现两次的字母
 */

// @lc code=start
func repeatedCharacter(s string) byte {
	mask := 0 // 掩码
	for _, c := range s {
		cur := 1 << (c - 'a') // 把每个字符向左移动
		if mask&cur > 0 {     // 如果>0代表意见存在 直接返回
			return byte(c)
		}
		// 否则代表不存在 加入到掩码中
		mask |= cur
	}
	return 0
}

// @lc code=end

