/*
 * @lc app=leetcode.cn id=799 lang=golang
 *
 * [799] 香槟塔
 */

// @lc code=start
func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := make([][]float64, query_row+2)
	for i := 0; i < query_row+2; i++ {
		dp[i] = make([]float64, i+1)
	}

	dp[0][0] = float64(poured)

	for i := 1; i <= query_row; i++ {
		for j := 0; j <= i; j++ {
			// 如果是第一列 只要它正上方的那一杯的 1/2的流量
			if j == 0 {
				// fmt.Printf("i:%d j:%d v:%f \n", i, j, max(0.0, dp[i-1][j]-1)/2)
				dp[i][j] = float64(max(0.0, dp[i-1][j]-1) / 2)
				continue
			}
			// 最后一列
			if j == i {
				// fmt.Printf("i:%d j:%d v:%f \n", i, j, max(0.0, dp[i-1][j-1]-1)/2)
				dp[i][j] = float64(max(0.0, dp[i-1][j-1]-1) / 2)
				continue
			}

			lastRow := max(0.0, (dp[i-1][j]-1)) / 2
			lastRowCol := max(0.0, (dp[i-1][j-1]-1)) / 2

			// 在中间的列
			dp[i][j] = lastRow + lastRowCol
		}
	}
	// fmt.Println(dp)
	return min(1.0, dp[query_row][query_glass])
}

func max(a, b float64) float64 {
	if a < b {
		return b
	}
	return a
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

