/*
 * @lc app=leetcode.cn id=1759 lang=golang
 *
 * [1759] 统计同构子字符串的数目
 */

// @lc code=start
func countHomogenous(s string) int {
	cnt, left, right, n := 0, 0, 0, len(s)
	const mod int = 1e9 + 7
	for ; left < n; left = right {
		// 初始右侧指针 从左侧开始移动
		right = left
		tmp := 0
		for right < n && s[left] == s[right] {
			tmp += right + 1 - left // 这里要先加
			// fmt.Printf("left:%d right:%d tmp:%d ch:%s \n", left, right, tmp, string(s[right]))
			right++
		}

		cnt += tmp
		cnt %= mod
		// fmt.Printf("left:%d right:%d cnt:%d subStr:%s\n", left, right, cnt, s[left:right])
	}
	return cnt
}

// @lc code=end

