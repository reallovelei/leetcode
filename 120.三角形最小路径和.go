/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	// 最后一行 当作初始值
	mini := triangle[len(triangle)-1]

	// 从倒数第二行开始，由下向上
	for row := len(triangle) - 2; row >= 0; row-- {
		// 列 从左到右
		for col := 0; col < len(triangle[row]); col++ {
			mini[col] = triangle[row][col] + min(mini[col], mini[col+1])
		}
	}
	return mini[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

