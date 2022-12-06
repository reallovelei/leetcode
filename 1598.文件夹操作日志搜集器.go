/*
 * @lc app=leetcode.cn id=1598 lang=golang
 *
 * [1598] 文件夹操作日志搜集器
 */

// @lc code=start
func minOperations(logs []string) int {
	depth := 0
	for _, l := range logs {
		if l == "./" {
			continue
		}

		if l == "../" {
			if depth > 0 {
				depth--
			}
			continue
		}

		depth++
	}
	return depth

}

// @lc code=end

