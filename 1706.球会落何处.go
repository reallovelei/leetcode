/*
 * @lc app=leetcode.cn id=1706 lang=golang
 *
 * [1706] 球会落何处
 1. 从第一列的球 开始往下每一行计算向左还是向右。
 2. 判断是否出界当 (列< 0 || 列 >= 列数)，或者卡住(当格子的方向 和 它同向相邻的格子的方向 不一样)
 3. 如果出界或者被卡住则 退出当前行的遍历， -1
*/

// @lc code=start
func findBall(grid [][]int) []int {
	// 第一行的 数量就是列数
	cols := len(grid[0])
	ans := make([]int, cols)

	//1. 从第一列的球 开始往下每一行计算向左还是向右。
	for c := 0; c < cols; c++ {
		col := c
		for _, row := range grid {
			// 当前方向
			direction := row[col]
			col = col + direction
			// 这一步有可能直接越界了 所以不能要
			// nextDirection := row[col]

			if col < 0 || col >= cols || direction != row[col] {
				col = -1
				break
			}

		}
		ans[c] = col
	}
	return ans
}

// @lc code=end

