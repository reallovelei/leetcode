/*
 * @lc app=leetcode.cn id=953 lang=golang
 *
 * [953] 验证外星语词典
 */

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	charIndex := [26]int{}
	for i, c := range order {
		charIndex[c-'a'] = i
	}

next:
	// 当前串 和 前一个串比较 所以从下标1开始
	for i := 1; i < len(words); i++ {

		for j := 0; j < len(words[i-1]) && j < len(words[i]); j++ {
			prev, cur := charIndex[words[i-1][j]-'a'], charIndex[words[i][j]-'a']
			if prev > cur {
				return false
			}
			// 不是相同字符 且字典序 前面 < 后面 下一轮
			if prev < cur {
				continue next
			}
		}
		// 如果前一个串的长度>当前串，且前部分都相等， 直接false
		if words[i-1] > words[i] {
			return false
		}

	}
	return true
}

// @lc code=end

